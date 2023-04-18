package api

const (
	TYPE_CHAT int32 = iota + 1
	TYPE_EDITS
)

type App struct {
	Name string `json:"name,omitempty" example:"商城商品"`
	Ord  int    `json:"ord,omitempty" example:"1"`
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
	SDParam     string `json:"sd_param" example:"{\"width\":512,\"height\":512}"`
	Tips        string `json:"tips" example:"为了获得更优的反馈，请描述商品名称与商品特性，如'YSL/圣罗兰 416； 红黑管唇釉，易上色，湿润'"`
	Ord         int    `json:"ord,omitempty" example:"1"`
}

type ChatAttr struct {
	ID      uint                    `json:"id,omitempty"`
	Type    int32                   `json:"type,omitempty" example:"1"`    // 1 纯文字 2 img
	Name    string                  `json:"name,omitempty" example:"商城商品"` // 属性名称
	Context []ChatCompletionMessage `json:"context,omitempty"`
	Tips    string                  `json:"tips" example:"为了获得更优的反馈，请描述商品名称与商品特性，如'YSL/圣罗兰 416； 红黑管唇釉，易上色，湿润'"`
	SDParam string                  `json:"sd_param" example:"{\"width\":512,\"height\":512}"`
	Ord     int                     `json:"ord,omitempty"`
}

type UserMessage struct {
	Content string
}

type ChatCompletionMessage struct {
	Role    string `json:"role" example:"user"`
	Content string `json:"content" example:"请问今天是几号?"`
}

type ImgMessage struct {
	Prompt string
	Negive string
}

type ChatCompletionResponse struct {
	Context string
}
