package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/repo/query"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// 生成图片
type SDController struct {
	// 这里可以注入一些服务或数据库连接
	Query *query.Query
}

func NewSDController() *SDController {
	return &SDController{
		Query: query.Use(infra.DB),
	}
}

type TextToImgMessage struct {
	UserMessage string // 用户输入
	SDParam
}

type SDParam struct {
	Prompt         string `json:"prompt,omitempty"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
}

// TextToImg swagger接口文档
// @Summary 文本转图片
// @Description 将文本转换为图片
// @Tags SD
// @Accept application/json
// @Produce jpeg
// @Param attr_id path int true "Attr ID" format(int32)
// @Param appData body TextToImgMessage true "更新的应用数据"
// @Success 200 {file} jpeg "图片文件"
// @Failure 500 {object} common.ErrorResponse "错误信息"
// @Router /v1/sd/{attr_id}/txt2img [post]
func (ctrl *SDController) TextToImg(c *gin.Context) {

	var param map[string]string = make(map[string]string)

	var x TextToImgMessage
	if err := c.BindJSON(&x); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if x.UserMessage != "" {
		var messages []openai.ChatCompletionMessage
		message := fmt.Sprintf(`
		"%s"
		###
		你是一个prompt工程师，请根据以上内容生成格式为 英文描述，图片风格随机
		`, x.UserMessage)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: message,
		})
		// 当 UserMessage 不为空，将获取 attr 属性翻译出结果
		resp, err := GetClient().CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		x.Prompt = x.Prompt + resp.Choices[0].Message.Content
	}

	param[ParamPrompt] = x.Prompt
	param[ParamNegativePrompt] = x.NegativePrompt

	var resp *http.Response
	resp, err := Request_Text2Img(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// 读取图片文件内容到内存中
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 设置响应头部的Content-Type和Content-Length字段
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(imageData)))

	// 将图片内容写入响应体
	c.Writer.Write(imageData)
}

var ImgUrl = func() string {
	return fmt.Sprintf("http://%s/", infra.Setting.SDHOST)
}

// ImgToImg 函数处理将一张图片文件上传并转换成另一张图片的请求。
// @Summary 图片转换
// @Tags SD
// @Description 将一张图片文件上传并转换成另一张图片
// @Accept multipart/form-data
// @Produce jpeg
// @Param attrID path int true "Attr ID"
// @Param image formData file true "待上传的图片文件"
// @Success 200 {file} jpeg "图片文件"
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/sd/{attr_id}/img2img [post]
func (ctrl *SDController) ImgToImg(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建一个文件缓冲区
	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)

	// 创建文件的表单字段
	fileField, err := fileWriter.CreateFormFile("file", "image.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 打开上传的文件
	srcFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer srcFile.Close()

	// 将上传的文件内容拷贝到表单字段
	_, err = io.Copy(fileField, srcFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 完成表单写入
	fileWriter.Close()

	url := ImgUrl() + "img2img"
	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, fileBuf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 创建 HTTP 客户端
	client := &http.Client{}
	// 设置 Content-Type 头部
	req.Header.Set("Content-Type", fileWriter.FormDataContentType())
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// 读取图片文件内容到内存中
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 设置响应头部的Content-Type和Content-Length字段
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(imageData)))

	// 将图片内容写入响应体
	c.Writer.Write(imageData)
}

const ParamPrompt string = "prompt"
const ParamNegativePrompt string = "negative_prompt"

// 需要手动关闭连接
func Request_Text2Img(param map[string]string) (*http.Response, error) {
	// 设置待发送的 prompt 和 negative_prompt 参数
	prompt := param[ParamPrompt]
	negativePrompt := param[ParamNegativePrompt]

	// 创建一个 buffer 用于构建请求体
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// 添加 prompt 参数到请求体
	_ = writer.WriteField("prompt", prompt)
	// 添加 negative_prompt 参数到请求体
	_ = writer.WriteField("negative_prompt", negativePrompt)

	// 关闭请求体写入
	writer.Close()

	URL := ImgUrl() + "txt2img"
	// 创建 POST 请求
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return nil, err
	}

	// 设置请求头，包含 Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return nil, err
	}
	return resp, nil
}

func init() {
	router := infra.GetApiEngine()
	chatCtrl := NewSDController()
	router.POST("/sd/:attr_id/txt2img", chatCtrl.TextToImg)
	router.POST("/sd/:attr_id/img2img", chatCtrl.ImgToImg)
}
