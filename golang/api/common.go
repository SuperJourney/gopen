package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SuperJourney/gopen/infra"
	"github.com/sashabaranov/go-openai"
)

const (
	TYPE_CHAT int32 = iota + 1
	TYPE_EDITS
)

type App struct {
	Name string `json:"name,omitempty" example:"商城商品"`
}

type UploadUrlResponse struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
}

type App_S struct {
	Name  string `json:"name,omitempty"`
	Attrs []Attr `json:"attrs,omitempty"`
}

type Attr struct {
	ID          uint   `json:"id" example:"0"`                                                              // Example ID
	Type        int32  `json:"type,omitempty" enums:"1,2" example:"1"`                                      // 1 chat completion 2 img
	Name        string `json:"name,omitempty" example:"商城商品"`                                               // Tab
	ContextType int32  `json:"context_type,omitempty" example:"1"`                                          // 1 chat 2 edit
	Context     string `json:"context,omitempty" example:"按照stable diffusion的关键字要求，给出主题鲜明的prompt,并使用英文文回答"` // 内容
}

type ChatAttr struct {
	ID      uint                    `json:"id,omitempty"`
	Type    int32                   `json:"type,omitempty" example:"1"`    // 1 纯文字 2 img
	Name    string                  `json:"name,omitempty" example:"商城商品"` // 属性名称
	Context []ChatCompletionMessage `json:"context,omitempty"`
}

type UserMessage struct {
	Content string
}

type ChatCompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ImgMessage struct {
	Prompt string
	Negive string
}

type ChatCompletionResponse struct {
	Context string
}

func GptEdits(msg string, instruction string) (string, error) {
	payload := fmt.Sprintf(`{
		"model": "text-davinci-edit-001",
		"input": "%s",
		"instruction": "%s"
	}`, msg, instruction)

	req, err := http.NewRequest("POST", infra.Setting.BaseURL+"/edits", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", infra.Setting.ApiToken))

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
