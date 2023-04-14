package api

import (
	"bytes"
	"context"
	"encoding/json"
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

	// 当 UserMessage 不为空，将获取 attr 属性翻译出结果
	// 读取图片文件内容到内存中
	imageData, shouldReturn := ctrl.textToImg(c)
	if shouldReturn {
		return
	}

	// 设置响应头部的Content-Type和Content-Length字段
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(imageData)))

	// 将图片内容写入响应体
	c.Writer.Write(imageData)
}

// @Summary 文本转图片链接
// @Description 将文本转换为图片链接
// @Tags SD
// @Accept application/json
// @Produce jpeg
// @Param attr_id path int true "Attr ID" format(int32)
// @Param appData body TextToImgMessage true "更新的应用数据"
// @Success 200 {object} UploadUrlResponse
// @Failure 500 {object} common.ErrorResponse "错误信息"
// @Router /v1/sd/{attr_id}/txt2imgurl [post]
func (ctrl *SDController) TextToImgUrl(c *gin.Context) {

	// 当 UserMessage 不为空，将获取 attr 属性翻译出结果
	// 读取图片文件内容到内存中
	imageData, shouldReturn := ctrl.textToImg(c)
	if shouldReturn {
		return
	}
	ctrl.toImageUrl(c, imageData)

}

func (*SDController) toImageUrl(c *gin.Context, imageData []byte) bool {
	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)

	fileField, err := fileWriter.CreateFormFile("file", "image.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return true
	}

	fileField.Write(imageData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return true
	}

	fileWriter.Close()

	url := infra.Setting.ImgUploadUrl

	req, err := http.NewRequest("POST", url, fileBuf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return true
	}

	client := &http.Client{}

	req.Header.Set("Content-Type", fileWriter.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return true
	}
	defer resp.Body.Close()

	var ret = &UploadUrlResponse{}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errRes openai.ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			reqErr := openai.RequestError{
				StatusCode: resp.StatusCode,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": reqErr.Error()})

		}
		errRes.Error.StatusCode = resp.StatusCode
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errRes.Error.Message})
	}
	json.NewDecoder(resp.Body).Decode(&ret)

	c.JSON(http.StatusOK, gin.H{
		"code": ret.Code,
		"data": ret.Data,
	})
	return false
}

type UploadUrlResponse struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
}

func (*SDController) textToImg(c *gin.Context) ([]byte, bool) {
	var param map[string]string = make(map[string]string)

	var x TextToImgMessage
	if err := c.BindJSON(&x); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
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

		resp, err := infra.GetClient().CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return nil, true
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
		return nil, true
	}
	defer resp.Body.Close()

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}
	return imageData, false
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
	imageData, shouldReturn := ctrl.imgToImg(c)
	if shouldReturn {
		return
	}

	// 设置响应头部的Content-Type和Content-Length字段
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Length", strconv.Itoa(len(imageData)))

	// 将图片内容写入响应体
	c.Writer.Write(imageData)
}

func (*SDController) imgToImg(c *gin.Context) ([]byte, bool) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}

	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)

	fileField, err := fileWriter.CreateFormFile("file", "image.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}

	srcFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}
	defer srcFile.Close()

	_, err = io.Copy(fileField, srcFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}

	fileWriter.Close()

	url := ImgUrl() + "img2img"

	req, err := http.NewRequest("POST", url, fileBuf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}

	client := &http.Client{}

	req.Header.Set("Content-Type", fileWriter.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}
	defer resp.Body.Close()

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, true
	}
	return imageData, false
}

// ImgToImg 函数处理将一张图片文件上传并转换成另一张图片的请求。
// @Summary 图片转换
// @Tags SD
// @Description 将一张图片文件上传并转换成另一张图片
// @Accept multipart/form-data
// @Produce jpeg
// @Param attrID path int true "Attr ID"
// @Param image formData file true "待上传的图片文件"
// @Success 200 {object} UploadUrlResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/sd/{attr_id}/img2imgurl [post]
func (ctrl *SDController) ImgToImgUrl(c *gin.Context) {
	imageData, shouldReturn := ctrl.imgToImg(c)
	if shouldReturn {
		return
	}
	ctrl.toImageUrl(c, imageData)
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

// TextToImg swagger接口文档
// @Summary 文本转图片
// @Description 将文本转换为图片
// @Tags SD
// @Accept application/json
// @Produce jpeg
// @Param attr_id path int true "Attr ID" format(int32)
// @Param appData body TextToImgMessage true "更新的应用数据"
// @Success 200 {object} UploadUrlResponse
// @Failure 500 {object} common.ErrorResponse "错误信息"
// @Router /v1/sd/{attr_id}/txt2imgurl [post]
func (ctrl *SDController) TextToImgURL(c *gin.Context) {
	ctrl.TextToImg(c)
}

func init() {
	router := infra.GetApiEngine()
	chatCtrl := NewSDController()
	router.POST("/sd/:attr_id/txt2img", chatCtrl.TextToImg)
	router.POST("/sd/:attr_id/img2img", chatCtrl.ImgToImg)

	router.POST("/sd/:attr_id/txt2imgurl", chatCtrl.TextToImgUrl)
	router.POST("/sd/:attr_id/img2imgurl", chatCtrl.ImgToImgUrl)
}
