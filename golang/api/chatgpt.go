package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SuperJourney/gopen/common"
	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/repo/model"
	"github.com/SuperJourney/gopen/repo/query"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

type ChatGptController struct {
	// 这里可以注入一些服务或数据库连接
	Query *query.Query
}

func NewChatGptController() *ChatGptController {
	return &ChatGptController{
		Query: query.Use(infra.DB),
	}
}

// Request generates chat completion text based on given input messages.
// @Summary Generate Chat Completion
// @Description Generate chat completion text based on input messages.
// @Tags ChatGpt
// @Accept json
// @Produce json
// @Param attrID path integer true "Attr ID"
// @Param userMessage body UserMessage true "User Messages"
// @Success 200 {object} ChatCompletionResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/gpt/{attrID}/chat-completion [post]
func (ctrl *ChatGptController) Request(c *gin.Context) {
	attrID, ok := GetAttrID(c)
	if !ok {
		return
	}

	var userMessage openai.ChatCompletionMessage
	if err := c.BindJSON(&userMessage); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	// 根据attrID 获取context
	db := ctrl.Query.Attr
	attrModel, err := db.Where(db.ID.Eq(uint(attrID))).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			common.Error(c, http.StatusBadRequest, errors.New("没有找到符合条件的attr"))
			return
		}
	}

	switch attrModel.ContextType {
	case int32(TYPE_EDITS):
		Edits(attrModel, c, userMessage.Content)
	default:
		// 创建文本生成请求
		// MaxTokens: int(infra.Setting.ChatGPT.MaxTokens),
		// Temperature: infra.Setting.ChatGPT.Temperature,
		// 调用OpenAI生成文本
		// 将生成的文本以JSON格式返回
		res, err := ChatCompletion(attrModel.Context, userMessage.Content)
		if err != nil {
			common.Error(c, http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{"context": res})
	}

}

func Edits(attrModel *model.Attr, c *gin.Context, msg string) {
	// 请求 URL 和参数
	// 创建 HTTP 请求
	// 设置请求头
	// 发送请求
	context, err := infra.GetClient().GptEdits(msg, attrModel.Context)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{"context": context})
	return
}

func ChatCompletion(preConversion string, msg string) (string, error) {
	var messages []openai.ChatCompletionMessage
	if err := json.Unmarshal([]byte(preConversion), &messages); err != nil {
		return "", err
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	}

	resp, err := infra.GetClient().CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func init() {
	router := infra.GetApiEngine()
	chatCtrl := NewChatGptController()
	router.POST("/gpt/:attr_id/chat-completion", chatCtrl.Request)
}
