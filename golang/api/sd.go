package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SuperJourney/gopen/common"
	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/repo/model"
	"github.com/SuperJourney/gopen/repo/query"
	"github.com/SuperJourney/gopen/utils"
	"github.com/SuperJourney/gopen/vars"
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
		Query: query.Use(vars.DB),
	}
}

type TextToImgMessage struct {
	UserMessage string `json:"user_message,omitempty"` // 用户输入
	SDParam     `json:"sd_param,omitempty"`
}

type SDParam struct {
	Prompt         string `json:"prompt,omitempty"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
	Width          int32  `json:"width,omitempty"`
	Height         int32  `json:"height,omitempty"`
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
	imageData, err := ctrl.textToImg(c)
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
	imageData, err := ctrl.textToImg(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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

	url := vars.Setting.ImgUploadUrl

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

func (ctrl *SDController) textToImg(c *gin.Context) ([]byte, error) {
	var x TextToImgMessage
	if err := c.BindJSON(&x); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	var prompt = c.PostForm("prompt")
	attrID, ok := GetAttrID(c)
	if !ok {
		return nil, errors.New("attr ID 为空")
	}
	db := ctrl.Query.Attr
	attrModel, err := db.Where(db.ID.Eq(uint(attrID))).First()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, errors.New("attr ID 为空"))
		return nil, err
	}
	usermessage := x.UserMessage

	prompt, err = zh2en(usermessage, prompt, attrModel, c)
	if err != nil {
		return nil, err
	}

	if attrModel.SDParam != "" {
		var sdparam SDParam
		if err := json.Unmarshal([]byte(attrModel.SDParam), &sdparam); err == nil {
			x.SDParam = sdparam
		}
	}

	x.SDParam.Prompt = x.SDParam.Prompt + prompt
	x.SDParam.NegativePrompt = x.NegativePrompt

	var resp *http.Response
	resp, err = Request_Text2Img(x.SDParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send request:", err)
		return nil, errors.New("textToImg server error")
	}

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	return imageData, nil
}

var ImgUrl = func() string {
	return fmt.Sprintf("http://%s/", vars.Setting.SDHOST)
}

// ImgToImg 函数处理将一张图片文件上传并转换成另一张图片的请求。
// @Summary 图片转换
// @Tags SD
// @Description 将一张图片文件上传并转换成另一张图片
// @Accept multipart/form-data
// @Produce jpeg
// @Param attr_id path int true "Attr ID"
// @Param image formData file true "待上传的图片文件"
// @Param use_message formData string true "用户提交内容"
// @Param width formData int false "生成图片宽度"
// @Param height formData int flase "生成图片高度"
// @Param prompt formData string flase "prompt"
// @Param negative_prompt formData string flase "negative_prompt"
// @Success 200 {file} jpeg "图片文件"
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/sd/{attr_id}/img2img [post]
func (ctrl *SDController) ImgToImg(c *gin.Context) {

	imageData, err := ctrl.imgToImg(c)
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

func (ctrl *SDController) imgToImg(c *gin.Context) ([]byte, error) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	usermessage := c.PostForm("user_message")
	var prompt = c.PostForm("prompt")
	attrID, ok := GetAttrID(c)
	if !ok {
		common.Error(c, http.StatusBadRequest, errors.New("attr ID 为空"))
		return nil, err
	}
	db := ctrl.Query.Attr
	attrModel, err := db.Where(db.ID.Eq(uint(attrID))).First()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, errors.New("attr ID 为空"))
		return nil, err
	}

	var sdParam = SDParam{
		Prompt:         c.PostForm("prompt"),
		NegativePrompt: c.PostForm("prompt_negative"),
	}

	width, err := strconv.ParseInt(c.PostForm("width"), 10, 32)
	if err == nil {
		sdParam.Width = int32(width)
	}

	height, err := strconv.ParseInt(c.PostForm("height"), 10, 32)
	if err == nil {
		sdParam.Height = int32(height)
	}

	if attrModel.SDParam != "" {
		var sdparamx SDParam
		if err := json.Unmarshal([]byte(attrModel.SDParam), &sdparamx); err == nil {
			sdParam = sdparamx
		}
	}

	prompt, err = zh2en(usermessage, prompt, attrModel, c)
	if err != nil {
		return nil, err
	}

	fileBuf := &bytes.Buffer{}
	fileWriter := multipart.NewWriter(fileBuf)
	sdParam.Prompt = sdParam.Prompt + prompt
	WriterAddParam(sdParam, fileWriter)

	fileField, err := fileWriter.CreateFormFile("file", "image.jpg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	srcFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	defer srcFile.Close()

	_, err = io.Copy(fileField, srcFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	fileWriter.Close()

	url := ImgUrl() + "img2img"

	req, err := http.NewRequest("POST", url, fileBuf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}

	client := &http.Client{}

	req.Header.Set("Content-Type", fileWriter.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	defer resp.Body.Close()

	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return nil, err
	}
	return imageData, nil
}

func zh2en(usermessage string, prompt string, attrModel *model.Attr, c *gin.Context) (string, error) {
	var err error
	if vars.Setting.BaiduTransate {
		chatprompt, err := utils.Translate(usermessage, "zh", "en")
		if err != nil {
			return "", err
		}
		prompt = prompt + chatprompt
	} else {
		if usermessage != "" {
			var chatprompt string
			switch attrModel.ContextType {
			case TYPE_CHAT:
				chatprompt, err = ChatCompletion(attrModel.Context, usermessage)
				if err != nil {
					return "", err
				}
			case TYPE_EDITS:
				chatprompt, err = vars.ChatClient.GptEdits(usermessage, attrModel.Context)
				if err != nil {
					return "", err
				}
			default:
				if err != nil {
					common.Error(c, http.StatusInternalServerError, errors.New("attr context_type invaild"))
					return "", err
				}
			}

			chatprompt, err := utils.Translate(chatprompt, "zh", "en")
			if err != nil {
				return "", err
			}

			prompt = prompt + chatprompt
		}
	}
	return prompt, nil
}

// ImgToImg 函数处理将一张图片文件上传并转换成另一张图片的请求。
// @Summary 图片转换
// @Tags SD
// @Description 将一张图片文件上传并转换成另一张图片
// @Accept multipart/form-data
// @Produce jpeg
// @Param attrID path int true "Attr ID"
// @Param image formData file true "待上传的图片文件"
// @Param use_message formData string true "用户提交内容"
// @Param width formData int false "生成图片宽度"
// @Param height formData int flase "生成图片高度"
// @Success 200 {object} UploadUrlResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/sd/{attr_id}/img2imgurl [post]
func (ctrl *SDController) ImgToImgUrl(c *gin.Context) {
	imageData, err := ctrl.imgToImg(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctrl.toImageUrl(c, imageData)
}

const ParamPrompt string = "prompt"
const ParamNegativePrompt string = "negative_prompt"
const ParamWidth = "width"
const ParamHeight = "height"

// 需要手动关闭连接
func Request_Text2Img(param SDParam) (*http.Response, error) {
	common.Info("text2img param: %v", param)
	// 创建一个 buffer 用于构建请求体
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// 设置待发送的 prompt 和 negative_prompt 参数
	// 添加 prompt 参数到请求体
	// 添加 negative_prompt 参数到请求体
	WriterAddParam(param, writer)

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
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send request:", err)
		return nil, errors.New("txt2img server error")
	}

	return resp, nil
}

func WriterAddParam(x SDParam, writer *multipart.Writer) {

	var param map[string]string = make(map[string]string)
	param[ParamPrompt] = x.Prompt
	param[ParamNegativePrompt] = x.NegativePrompt
	param[ParamHeight] = strconv.Itoa(int(x.Height))
	param[ParamWidth] = strconv.Itoa(int(x.Width))

	if prompt, ok := param[ParamPrompt]; ok {
		_ = writer.WriteField("prompt", prompt)
	}

	if negativePrompt, ok := param[ParamNegativePrompt]; ok {
		_ = writer.WriteField("negative_prompt", negativePrompt)
	}

	if width, ok := param[ParamWidth]; ok {
		if width != "" {
			_ = writer.WriteField("width", width)
		}

	}
	if height, ok := param[ParamHeight]; ok {
		if height != "" {
			_ = writer.WriteField("height", height)
		}
	}

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
	router.POST("/sd/:attr_id/txt2imgurl", chatCtrl.TextToImgUrl)

	router.POST("/sd/:attr_id/img2img", chatCtrl.ImgToImg)
	router.POST("/sd/:attr_id/img2imgurl", chatCtrl.ImgToImgUrl)
}
