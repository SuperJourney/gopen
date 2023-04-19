package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// 商品
type App struct {
	Name  string `json:"name,omitempty"`
	Attrs []Attr `json:"attrs,omitempty"`
	Ord   int    `json:"ord,omitempty"`
	gorm.Model
}

type Attr struct {
	Type        int32  `json:"type,omitempty"`                     // 1 chat completion 2 img
	Name        string `json:"name,omitempty"`                     // Tab
	ContextType int32  `json:"context_type,omitempty" example:"1"` // 1 chat 2 edit
	Context     string `json:"context,omitempty"`                  // 内容
	AppID       int32  `json:"app_id,omitempty"`
	SDParam     string `json:"sd_param"`
	Tips        string `json:"tips"`
	Ord         int    `json:"ord,omitempty"`
	gorm.Model
}

type Img struct {
	Prompt         string `json:"prompt,omitempty"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
}

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
