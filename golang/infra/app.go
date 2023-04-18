package infra

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

var Engine = gin.Default()

func GetApiEngine() *gin.RouterGroup {
	return Engine.Group("api/v1")
}

func GetApiEngineV2() *gin.RouterGroup {
	return Engine.Group("api/v2")
}

type ChatClient interface {
	CreateChatCompletion(
		ctx context.Context,
		request openai.ChatCompletionRequest,
	) (response openai.ChatCompletionResponse, err error)
}

func GetClient() *openai.Client {

	config := openai.DefaultConfig(Setting.ApiToken)
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

	config.BaseURL = Setting.BaseURL

	client := openai.NewClientWithConfig(config)
	return client
}
