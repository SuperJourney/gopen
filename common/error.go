package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccResponse struct {
	Succ string `json:"succ"`
}

func Error(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, err)
}

func Succ(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"succ": true})
}
