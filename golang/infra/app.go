package infra

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

	GptEdits(msg string, instruction string) (string, error)
}

type OpenAiChatClient struct {
	*openai.Client
}

func (a *OpenAiChatClient) GptEdits(msg string, instruction string) (string, error) {
	payload := fmt.Sprintf(`{
		"model": "text-davinci-edit-001",
		"input": "%s",
		"instruction": "%s"
	}`, msg, instruction)

	req, err := http.NewRequest("POST", Setting.BaseURL+"/edits", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Setting.ApiToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
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
			return "", &reqErr

		}
		errRes.Error.StatusCode = res.StatusCode
		return "", errRes.Error
	}

	var v *openai.EditsResponse = &openai.EditsResponse{}
	if err = json.NewDecoder(res.Body).Decode(v); err != nil {
		return "", err
	}
	return v.Choices[0].Text, nil
}

func GetClient() ChatClient {

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
	return &OpenAiChatClient{
		Client: client,
	}
}
