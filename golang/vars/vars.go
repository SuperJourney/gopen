package vars

import (
	"github.com/SuperJourney/gopen/config"
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ChatClientIFace interface {
	ChatCompletion(request openai.ChatCompletionRequest) (string, error)
	GptEdits(msg string, instruction string) (string, error)
}

var ChatClient ChatClientIFace

var Setting *config.Setting

var DB *gorm.DB

func init() {
	Setting = config.LoadConfig()

	var err error
	DB, err = gorm.Open(sqlite.Open(Setting.DBFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
