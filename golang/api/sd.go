package api

// import (
// 	"net/http"

// 	"github.com/SuperJourney/gopen/infra"
// 	"github.com/gin-gonic/gin"
// )

// // 生成图片
// type SDController struct {
// }

// func NewSDController() *SDController {
// 	return &SDController{}
// }

// type SDTextMsg struct {
// 	Prompt         string
// 	NegativePrompt string
// }

// var host = "popoen:"

// func (ctrl *SDController) TextToImg(c *gin.Context) {

// }

// func (ctrl *SDController) ImgToImg(c *gin.Context) {
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// }

// func init() {
// 	router := infra.GetApiEngine()
// 	chatCtrl := NewSDController()
// 	router.POST("/gpt/:attr_id/chat-completion", chatCtrl.TextToImg)
// }
