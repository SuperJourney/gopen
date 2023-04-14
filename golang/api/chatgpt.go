package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
		ChatCompletion(attrModel, c, userMessage.Content)
	}

}

func Edits(attrModel *model.Attr, c *gin.Context, msg string) {
	// s := openai.GPT3TextDavinci001
	// req := openai.EditsRequest{
	// 	Model:       s,
	// 	Input:       msg,
	// 	Instruction: attrModel.Context,
	// }

	// resp, err := infra.GetClient().Edits()
	// ts(context.Background(), req)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return true
	// }

	// 请求 URL 和参数
	payload := fmt.Sprintf(`{
		"model": "text-davinci-edit-001",
		"input": "%s",
		"instruction": "%s"
	}`, msg, attrModel.Context)

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", infra.Setting.BaseURL+"/edits", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", infra.Setting.ApiToken))

	// 发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes openai.ErrorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err != nil || errRes.Error == nil {
			reqErr := openai.RequestError{
				StatusCode: res.StatusCode,
				Err:        err,
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": reqErr.Error()})

		}
		errRes.Error.StatusCode = res.StatusCode
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errRes.Error.Message})
	}

	var v *openai.EditsResponse = &openai.EditsResponse{}
	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"context": v.Choices[0].Text})

}

func ChatCompletion(attrModel *model.Attr, c *gin.Context, msg string) bool {
	var messages []openai.ChatCompletionMessage
	if err := json.Unmarshal([]byte(attrModel.Context), &messages); err != nil {
		common.Info("err:%v", err)
		common.Error(c, http.StatusInternalServerError, errors.New("attr信息异常"))
		return true
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return true
	}

	c.JSON(http.StatusOK, gin.H{"context": resp.Choices[0].Message.Content})
	return false
}

func init() {
	router := infra.GetApiEngine()
	chatCtrl := NewChatGptController()
	router.POST("/gpt/:attr_id/chat-completion", chatCtrl.Request)
}
