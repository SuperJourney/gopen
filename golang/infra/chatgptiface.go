package infra

import (
	"github.com/SuperJourney/gopen/infra/chatgpt"
	"github.com/SuperJourney/gopen/vars"
)

func init() {

	vars.ChatClient = chatgpt.GetClient()

}
