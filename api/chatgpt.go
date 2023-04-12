package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/SuperJourney/gopen/common"
	"github.com/SuperJourney/gopen/infra"
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

func GetClient() *openai.Client {

	config := openai.DefaultConfig(infra.Setting.ApiToken)
	// proxyUrl, err := url.Parse("http://143.42.236.71:8443/")
	// if err != nil {
	// 	panic(err)
	// }
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyUrl),
	// }
	// config.HTTPClient = &http.Client{
	// 	Transport: transport,
	// }

	config.BaseURL = "http://143.42.236.71:8443/v1"

	client := openai.NewClientWithConfig(config)
	return client
}

// @Summary Generate Text Completion
// @Description Generate text completion based on prompt
// @Tags v1
// @Accept json
// @Produce json
// @Param prompt query string true "Prompt for text completion"
// @Param max_tokens query int false "Maximum number of tokens to generate"
// @Success 200 {string} string "Generated text"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/gpt/text-completion [get]
func TextCompletion(c *gin.Context) {
	// 获取请求参数
	prompt := c.Query("prompt")
	maxTokens := infra.Setting.ChatGPT.MaxTokens // 默认最大token数为100

	// 创建文本生成请求
	req := openai.CompletionRequest{
		Model:     infra.Setting.ChatGPT.Model,
		Prompt:    prompt,
		MaxTokens: int(maxTokens),
	}

	// 调用OpenAI生成文本
	resp, err := GetClient().CreateCompletion(context.Background(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将生成的文本以JSON格式返回
	c.JSON(http.StatusOK, resp.Choices[0].Text)
}

// ChatCompletion generates chat completion text based on given input messages.
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
func (ctrl *ChatGptController) ChatCompletion(c *gin.Context) {
	attrID, ok := GetAttrID(c)
	if !ok {
		return
	}

	var userMessage openai.ChatCompletionMessage
	if err := c.BindJSON(&userMessage); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}
	userMessage.Role = openai.ChatMessageRoleUser

	// 根据attrID 获取context
	db := ctrl.Query.Attr
	attrModel, err := db.Where(db.ID.Eq(uint(attrID)), db.Type.Eq(TYPE_CHAT_COMPLETION)).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			common.Error(c, http.StatusBadRequest, errors.New("没有找到符合条件的attr"))
			return
		}
	}
	var messages []openai.ChatCompletionMessage
	if err := json.Unmarshal([]byte(attrModel.Context), &messages); err != nil {
		common.Info("err:%v", err)
		common.Error(c, http.StatusInternalServerError, errors.New("attr信息异常"))
		return
	}

	messages = append(messages, userMessage)

	// Example
	// []openai.ChatCompletionMessage{
	// 	{
	// 		Role:    openai.ChatMessageRoleSystem,
	// 		Content: "你是一个商品设计小组手,根据用户少量的提示词就可以给出极具创造力，吸引力的标题",
	// 	},
	// 	{
	// 		Role:    openai.ChatMessageRoleUser,
	// 		Content: "口红，女性， 颜色",
	// 	},
	// 	{
	// 		Role:    openai.ChatMessageRoleAssistant,
	// 		Content: "最时尚的口红色彩",
	// 	},
	// 	{
	// 		Role:    openai.ChatMessageRoleUser,
	// 		Content: "球鞋，跳高，特步",
	// 	},
	// }
	// 创建文本生成请求
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
		// MaxTokens: int(infra.Setting.ChatGPT.MaxTokens),
		// Temperature: infra.Setting.ChatGPT.Temperature,
	}

	// 调用OpenAI生成文本
	resp, err := GetClient().CreateChatCompletion(context.Background(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将生成的文本以JSON格式返回
	c.JSON(http.StatusOK, gin.H{"context": resp.Choices[0].Message.Content})
}

func init() {
	router := infra.GetApiEngine()
	chatCtrl := NewChatGptController()
	router.POST("/gpt/:attr_id/chat-completion", chatCtrl.ChatCompletion)
}
