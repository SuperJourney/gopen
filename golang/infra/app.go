package infra

import (
	"github.com/gin-gonic/gin"
)

var Engine = gin.Default()

func GetApiEngine() *gin.RouterGroup {
	return Engine.Group("api/v1")
}

func GetApiEngineV2() *gin.RouterGroup {
	return Engine.Group("api/v2")
}
