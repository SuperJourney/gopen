package api

import (
	"context"
	"net/http"

	"github.com/SuperJourney/gopen/infra"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

var client = GetClient()

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
	resp, err := client.CreateCompletion(context.Background(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将生成的文本以JSON格式返回
	c.JSON(http.StatusOK, resp.Choices[0].Text)
}

// @Summary 使用OpenAI生成对话文本
// @Description 基于OpenAI的Chat Completion API，生成对话文本。
// @Tags v1
// @Accept json
// @Produce json
// @Param role query string true "对话角色"
// @Param content query string true "对话内容"
// @Success 200 {string} string "成功生成对话文本"
// @Failure 500 {object} error "生成对话文本失败"
// @Router /v1/gpt/chat-completion [get]
func ChatCompletion(c *gin.Context) {

	// 创建文本生成请求
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "你是一个商品设计小组手,根据用户少量的提示词就可以给出极具创造力，吸引力的标题",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "口红，女性， 颜色",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "最时尚的口红色彩",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "球鞋，跳高，特步",
			},
		},
		// MaxTokens: int(infra.Setting.ChatGPT.MaxTokens),
		// Temperature: infra.Setting.ChatGPT.Temperature,
	}

	// 调用OpenAI生成文本
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将生成的文本以JSON格式返回
	c.JSON(http.StatusOK, resp.Choices[0].Message.Content)
}

func init() {
	r := infra.GetApiEngine()
	r.GET("/gpt/text-completion", TextCompletion)
	r.GET("/gpt/chat-completion", ChatCompletion)
}
