package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/SuperJourney/gopen/common"
	"github.com/SuperJourney/gopen/infra"
	"github.com/SuperJourney/gopen/repo/model"
	"github.com/SuperJourney/gopen/repo/query"
	"github.com/SuperJourney/gopen/vars"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// type Attr struct {
// 	ID      uint   `gorm:"primarykey"`
// 	Type    int32  `json:"type,omitempty"`    // 1 chat completion 2 img
// 	Name    string `json:"name,omitempty"`    // Tab
// 	Context string `json:"context,omitempty"` // 内容
// }

// type App_S struct {
// 	App
// 	Attrs []Attr
// }

type AppController struct {
	// 这里可以注入一些服务或数据库连接
	Query *query.Query
}

func NewAppController() *AppController {
	return &AppController{
		Query: query.Use(vars.DB),
	}
}

// @Summary Get all apps
// @Description Get all apps
// @Tags App
// @Produce json
// @Success 200 {array} App
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps [get]
func (ctrl *AppController) GetApps(c *gin.Context) {
	apps, err := ctrl.Query.App.Order(ctrl.Query.App.Ord).Find()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"apps": apps,
	})
}

// @Summary Get app by ID
// @Description Get app by ID
// @Tags App
// @Produce json
// @Param id path int true "App ID"
// @Success 200 {object} App_S
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{id} [get]
func (ctrl *AppController) GetApp(c *gin.Context) {
	appID := c.Param("app_id")

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
	appDB := ctrl.Query.App
	app, err := appDB.Where(appDB.ID.Eq(uint(id))).First()
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

	// Query the app by ID
	attrDB := ctrl.Query.Attr
	attrs, err := attrDB.Where(attrDB.AppID.Eq(int32(id))).Find()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	for _, v := range attrs {
		app.Attrs = append(app.Attrs, *v)
	}

	c.JSON(http.StatusOK, gin.H{
		"app": app,
	})
}

// @Summary Create a new app
// @Description Creates a new app with the provided data
// @Tags App
// @Accept json
// @Produce json
// @Param app body App true "App data"
// @Success 200 {object} App
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
	if err := db.Create(&app); err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"app": app,
	})
}

// @Summary 更新应用
// @Description 根据提供的应用ID和更新的应用数据更新应用信息
// @Tags App
// @Accept json
// @Produce json
// @Param app_id path string true "应用ID"
// @Param appData body App true "更新的应用数据"
// @Success 200 {object} App
// @Failure 400 {object} common.ErrorResponse
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{id} [put]
func (ctrl *AppController) UpdateApp(c *gin.Context) {
	// Get the app ID from the URL parameter
	appID := c.Param("app_id")

	// Parse the request body to get the updated app data
	var updatedApp *model.App // Replace "App" with the appropriate struct for your app
	err := c.ShouldBindJSON(&updatedApp)
	if err != nil {
		// Handle error if parsing fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	appDB := ctrl.Query.App
	// Fetch the existing app from the database
	existingApp, err := appDB.Where(appDB.ID.Eq(uint(id))).First()
	if err != nil {
		// Handle error if fetching app fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the existing app with the updated app data
	existingApp.Name = updatedApp.Name
	// Update other fields as needed
	existingApp.Ord = updatedApp.Ord

	// Save the updated app to the database
	appDB.Create(existingApp)

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"app": existingApp,
	})
}

// @Summary Delete an app
// @Description Delete an app by app ID
// @Tags App
// @Accept json
// @Produce json
// @Param app_id path int true "App ID"
// @Success 200 {object} common.SuccResponse "Successfully deleted app"
// @Failure 400 {object} common.ErrorResponse "Invalid app ID"
// @Failure 500 {object} common.ErrorResponse "Internal server error"
// @Router /v1/apps/{app_id} [delete]
func (ctrl *AppController) DeleteApp(c *gin.Context) {
	// Get the app ID from the URL parameter
	appID := c.Param("app_id")

	// Convert the app ID to an integer
	id, err := strconv.Atoi(appID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid app ID",
		})
		return
	}

	appDB := ctrl.Query.App
	// Fetch the existing app from the database
	existingApp, err := appDB.Where(appDB.ID.Eq(uint(id))).First()
	if err != nil {
		// Handle error if fetching app fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	appDB.Delete(existingApp)
	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

func init() {
	router := infra.GetApiEngine()
	appCtrl := NewAppController()

	// App 相关路由
	router.GET("/apps", appCtrl.GetApps)
	router.GET("/apps/:app_id", appCtrl.GetApp)
	router.POST("/apps", appCtrl.CreateApp)
	router.PUT("/apps/:app_id", appCtrl.UpdateApp)
	router.DELETE("/apps/:app_id", appCtrl.DeleteApp)

}
