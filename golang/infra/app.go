package infra

import (
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

// func ss() {
// 	stream, err := GetClient().CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{})
// 	if err != nil {
// 		fmt.Printf("CompletionStream error: %v\n", err)
// 		return
// 	}
// 	defer stream.Close()

// 	for {
// 		response, err := stream.Recv()
// 		if errors.Is(err, io.EOF) {
// 			fmt.Println("Stream finished")
// 			return
// 		}

// 		if err != nil {
// 			fmt.Printf("Stream error: %v\n", err)
// 			return
// 		}

// 		fmt.Printf("Stream response: %v\n", response)
// 	}
// }

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
