package config

import (
	"github.com/pelletier/go-toml"
)

var ConfigPath = "/workspaces/gopen/config.toml"

type Setting struct {
	ChatGPT
	AppSetting
}

type AppSetting struct {
	ConfigFile    string `desc:"The path to the configuration file to load"`
	DBFile        string ``
	SDHOST        string
	ImgUploadUrl  string
	BaiduTransate bool
	BaiduAppId    string
	BaiduApiKey   string
}

type ChatGPT struct {
	ProxyUrl    string `desc:"The URL of the proxy server to use for requests"`
	ApiToken    string `desc:"The API token to authenticate requests to the GPT service" flag:"api-token"`
	MaxTokens   int32  `desc:"The maximum number of tokens to generate for each response"`
	Model       string `desc:"The name or path of the GPT model to use"`
	Temperature float32
	BaseURL     string
}

func LoadConfig() *Setting {
	tree, err := toml.LoadFile(ConfigPath)
	if err != nil {
		panic(err)
	}

	// Unmarshal the TOML data into a Setting struct.
	var setting Setting
	if err := tree.Unmarshal(&setting); err != nil {
		panic(err)
	}

	return &setting
}
