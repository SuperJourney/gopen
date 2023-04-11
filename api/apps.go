package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SuperJourney/gopen/common"
	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/repo/model"
	"github.com/SuperJourney/gopen/repo/query"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppController struct {
	// 这里可以注入一些服务或数据库连接
	Query *query.Query
}

func NewAppController() *AppController {
	return &AppController{
		Query: query.Use(infra.DB),
	}
}

// @Summary Get all apps
// @Description Get all apps
// @Tags v1
// @Produce json
// @Success 200 {array} model.App
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps [get]
func (ctrl *AppController) GetApps(c *gin.Context) {
	apps, err := ctrl.Query.App.Find()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"apps": apps,
	})
}

// @Summary Get app by ID
// @Description Get app by ID
// @Tags v1
// @Produce json
// @Param id path int true "App ID"
// @Success 200 {object} model.App
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{id} [get]
func (ctrl *AppController) GetApp(c *gin.Context) {
	appID := c.Param("id")

	// Validate the app ID parameter
	if appID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing app ID",
		})
		return
	}

	// Convert the app ID to an integer
	id, err := strconv.Atoi(appID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid app ID",
		})
		return
	}

	// Query the app by ID
	db := ctrl.Query.App
	app, err := db.Where(db.ID.Eq(uint(id))).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "app not found",
			})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	// Return the app as JSON
	c.JSON(http.StatusOK, gin.H{
		"app": app,
	})
}

// @Summary Create a new app
// @Description Creates a new app with the provided data
// @Tags v1
// @Accept json
// @Produce json
// @Param app body model.App true "App data"
// @Success 200 {object} model.App
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps [post]
func (ctrl *AppController) CreateApp(c *gin.Context) {
	app := model.App{}

	// Parse request body to extract app data
	err := c.BindJSON(&app)
	if err != nil {
		common.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Validate app data
	if app.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "App name is required"})
		return
	}
	db := ctrl.Query.App
	// Save app data to database
	db.Create(&app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save app data"})
		return
	}

	c.JSON(http.StatusOK, app)
}

func (ctrl *AppController) UpdateApp(c *gin.Context) {

}

func (ctrl *AppController) DeleteApp(c *gin.Context) {
	// 删除 App
}

func (ctrl *AppController) GetAttrs(ct *gin.Context) {

}

func init() {
	router := infra.GetApiEngine()
	appCtrl := NewAppController()

	// App 相关路由
	router.GET("/apps", appCtrl.GetApps)
	router.GET("/apps/:id", appCtrl.GetApp)
	router.POST("/apps", appCtrl.CreateApp)
	router.PUT("/apps/:id", appCtrl.UpdateApp)
	router.DELETE("/apps/:id", appCtrl.DeleteApp)

}
