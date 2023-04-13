package api

var (
	TYPE_CHAT_COMPLETION int32 = 1
)

type App struct {
	Name string `json:"name,omitempty" example:"商城商品"`
}

type Attr struct {
	ID      uint   `json:"id" example:"0"`                                                             // Example ID
	Type    int32  `json:"type,omitempty" enums:"1,2" example:"1"`                                     // 1 chat completion 2 img
	Name    string `json:"name,omitempty" example:"商城商品"`                                              // Tab
	Context string `json:"context,omitempty" example:"{\"content\":\"请问当前时间是几点？\",\"role\":\"user\"}"` // 内容
}

type ChatAttr struct {
	ID      uint
	Type    int32  `json:"type,omitempty" example:"1"`    // 1 chat completion 2 img
	Name    string `json:"name,omitempty" example:"商城商品"` // Tab
	Context []ChatCompletionMessage
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
