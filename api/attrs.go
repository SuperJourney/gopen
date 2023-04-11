package api

// import (
// 	"errors"
// 	"net/http"
// 	"strconv"

// 	"github.com/SuperJourney/gopen/common"
// 	"github.com/SuperJourney/gopen/infra"
// 	"github.com/SuperJourney/gopen/repo/model"
// 	"github.com/SuperJourney/gopen/repo/query"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type AttrsController struct {
// 	// 这里可以注入一些服务或数据库连接
// 	Query *query.Query
// }

// func NewAttrsController() *AttrsController {
// 	return &AttrsController{
// 		Query: query.Use(infra.DB),
// 	}
// }

// // @Summary Get all apps
// // @Description Get all apps
// // @Tags v1
// // @Produce json
// // @Success 200 {array} model.Attr
// // @Failure 500 {object} common.ErrorResponse
// // @Router /v1/apps [get]
// func (ctrl *AttrsController) GetAttrs(c *gin.Context) {
// 	apps, err := ctrl.Query.Attr.Find()
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"apps": apps,
// 	})
// }

// // @Summary Get app by ID
// // @Description Get app by ID
// // @Tags v1
// // @Produce json
// // @Param id path int true "Attrs ID"
// // @Success 200 {object} model.Attr
// // @Failure 400 {object} common.ErrorResponse
// // @Failure 404 {object} common.ErrorResponse
// // @Failure 500 {object} common.ErrorResponse
// // @Router /v1/apps/{id} [get]
// func (ctrl *AttrsController) GetAttr(c *gin.Context) {
// 	appID := c.Param("id")

// 	// Validate the app ID parameter
// 	if appID == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "missing app ID",
// 		})
// 		return
// 	}

// 	// Convert the app ID to an integer
// 	id, err := strconv.Atoi(appID)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid app ID",
// 		})
// 		return
// 	}

// 	// Query the app by ID
// 	db := ctrl.Query.Attr
// 	app, err := db.Where(db.ID.Eq(uint(id))).First()
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error": "app not found",
// 			})
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 		}
// 		return
// 	}

// 	// Return the app as JSON
// 	c.JSON(http.StatusOK, gin.H{
// 		"app": app,
// 	})
// }

// // @Summary Create a new app
// // @Description Creates a new app with the provided data
// // @Tags v1
// // @Accept json
// // @Produce json
// // @Param app body model.Attr true "Attrs data"
// // @Success 200 {object} model.Attr
// // @Failure 400 {object} common.ErrorResponse
// // @Failure 500 {object} common.ErrorResponse
// // @Router /v1/apps [post]
// func (ctrl *AttrsController) CreateAttr(c *gin.Context) {
// 	app := model.Attr{}

// 	// Parse request body to extract app data
// 	err := c.BindJSON(&app)
// 	if err != nil {
// 		common.Info(err.Error())
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
// 		return
// 	}

// 	// Validate app data
// 	if app.Name == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Attrs name is required"})
// 		return
// 	}
// 	db := ctrl.Query.Attr
// 	// Save app data to database
// 	db.Create(&app)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save app data"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, app)
// }

// func (ctrl *AttrsController) UpdateAttr(c *gin.Context) {

// }

// func (ctrl *AttrsController) DeleteAttr(c *gin.Context) {
// 	// 删除 Attrs
// }

// func init() {
// 	router := infra.GetApiEngine()
// 	attrsCtrl := NewAttrsController()

// 	// Prompt 相关路由
// 	router.GET("/apps/:app_id/attrs", attrsCtrl.GetAttrs)
// 	router.GET("/apps/:app_id/attrs/:attr", attrsCtrl.GetAttr)

// 	// 根据类型创建
// 	router.POST("/apps/:app_id/:type/attrs/", attrsCtrl.CreateAttr)
// 	router.PUT("/apps/:app_id/:type/attrs/:id", attrsCtrl.UpdateAttr)

// 	router.DELETE("/apps/:app_id/attrs/:id", attrsCtrl.DeleteAttr)

// }
