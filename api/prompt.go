package api

import (
	"net/http"

	"github.com/SuperJourney/gopen/infra"
	"github.com/gin-gonic/gin"
)

type Prompt struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var prompts = []Prompt{
	{ID: "1", Text: "Enter your name"},
	{ID: "2", Text: "What's your favorite color?"},
}

// @Summary Retrieve a prompt by ID
// @Description Retrieves a prompt by ID
// @Produce json
// @Tags v1
// @Param id path string true "Prompt ID"
// @Success 200 {object} Prompt "Successful operation"
// @Router /v1/prompt/{id} [get]
func GetPrompt(c *gin.Context) {
	id := c.Param("id")
	for _, prompt := range prompts {
		if prompt.ID == id {
			c.JSON(http.StatusOK, prompt)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
}

// @Summary Retrieve a prompt by ID
// @Description Retrieves a prompt by ID
// @Produce json
// @Tags v2
// @Param id path string true "Prompt ID"
// @Success 200 {object} Prompt "Successful operation"
// @Router /v2/prompt/{id} [get]
func GetPromptV2(c *gin.Context) {
	id := c.Param("id")
	for _, prompt := range prompts {
		if prompt.ID == id {
			c.JSON(http.StatusOK, prompt)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
}

func GetPrompts(c *gin.Context) {
	c.JSON(http.StatusOK, prompts)
}

func CreatePrompt(c *gin.Context) {
	var prompt Prompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prompt.ID = string(len(prompts) + 1)
	prompts = append(prompts, prompt)
	c.JSON(http.StatusCreated, prompt)
}

func DeletePrompt(c *gin.Context) {
	id := c.Param("id")
	for i, prompt := range prompts {
		if prompt.ID == id {
			prompts = append(prompts[:i], prompts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Prompt deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Prompt not found"})
}

func init() {
	r := infra.GetApiEngine()
	r.GET("/prompt/:id", GetPrompt)
	r.GET("/prompt", GetPrompts)
	r.POST("/prompt", CreatePrompt)
	r.DELETE("/prompt/:id", DeletePrompt)
}
