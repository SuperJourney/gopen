package model

import (
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// 商品

type App struct {
	Name  string `json:"name,omitempty"`
	Attrs []Attr `json:"attrs,omitempty"`
	BaseModel
}

type Attr struct {
	Type    int32  `json:"type,omitempty"`    // 1 chat completion 2 img
	Name    string `json:"name,omitempty"`    // Tab
	Context string `json:"context,omitempty"` // 内容
	AppID   int32  `json:"app_id,omitempty"`
	BaseModel
}

type BaseModel struct {
	ID        uint  `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

// type AppAttr struct {
// 	AppId  int32
// 	AttrId int32
// }

var contextChat = []openai.ChatCompletionMessage{
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
}

type Img struct {
	Prompt         string `json:"prompt,omitempty"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
}

// type Prompt struct {
// 	ID     string `json:"id"`
// 	Type   int32  //
// 	Scenes string
// 	Format int32    //查找是不是又 {} 占位符，否者直接放在最后面  == 0 特殊，不使用任何prompt
// 	Prompt []string `json:"text"`
// 	// Input  string
// }

// type Format struct {
// 	ID      int32
// 	Name    string // 根据
// 	Context string //请根据相关提示进行联想，$1  给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容 {}
// 	Example string
// }

// // 场景值 ， 图片风格，等
// type Scene struct {
// 	ID    int32
// 	Scene string
// }

// 输入：
// tempalte : “请根据相关提示进行联想，$0  给出指定$1 ， $1 长度 $2 个字，不要包含除了 $1 的其他内容 ”,
// context: [ "标题"，“4-10”]
// input: "口红"

// 输出： “请根据相关提示进行联想，口红  给出指定标题 ， 标题 长度 4-10 个字，不要包含除了 标题的其他内容”
// 请用golang函数实现

func Replace() {

	// $0 代表用户输入
	template := "请根据相关提示进行联想，$0 给出指定$1，$1长度$2个字，不要包含除了$1的其他内容"

	context := []string{"标题", "4-10"}
	input := "口红"
	context = append([]string{input}, context...)

	for i, val := range context {
		placeholder := fmt.Sprintf("$%d", i)
		template = strings.Replace(template, placeholder, val, -1)
	}

	fmt.Println(template)
}

// 文本

// -- 请根据指定关键字进行联想，口红,特征： 女性，红色，给出指定标题， 标题长度5-10个字， 不要包含除了标题的其他内容

// 请根据相关提示进行联想，$1  给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容

// 请根据相关提示进行联想，口红 给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容

// 类型：标题
// prompt拼接：

// type : 1
// 请根据相关提示进行联系，$1  给出指定 $2 ， $2 长度 $3 个字，不要包含除了 $2 的其他内容

// 图片

// 风格1 风格2 风格3 + prompt

// 顶层结构体： 应用
// 应用下可以添加标题描述，或者 图片描述
// 所以我定义了一个Context代表图片和标题

// 请使用golang给出合适的结构体定义
