package api

import (
	"encoding/json"
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

type AttrController struct {
	// 这里可以注入一些服务或数据库连接
	Query *query.Query
}

func NewAttrController() *AttrController {
	return &AttrController{
		Query: query.Use(infra.DB),
	}
}

// @Summary Get all Attrs
// @Description Get all attrs
// @Tags Attr
// @Produce json
// @Param app_id path int true "App ID"
// @Success 200 {array} Attr
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{app_id}/attrs [get]
func (ctrl *AttrController) GetAttrs(c *gin.Context) {
	appID := c.Param("app_id")

	// Validate the app ID parameter
	if appID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing app ID",
		})
		return
	}

	// Convert the app ID to an integer
	appIdIntger, err := strconv.Atoi(appID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid app ID",
		})
		return
	}

	// Query the app by ID
	attrDB := ctrl.Query.Attr
	attrs, err := attrDB.Where(attrDB.AppID.Eq(int32(appIdIntger))).Debug().Find()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"attrs": attrs,
	})
}

// @Summary Get Attr by ID
// @Description Get Attr by ID
// @Tags Attr
// @Produce json
// @Param app_id path int true "App ID"
// @Param attr_id path int true "Attr ID"
// @Success 200 {object} Attr
// @Failure 400 {object} common.ErrorResponse
// @Failure 404 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{app_id}/attrs/{attr_id} [get]
func (ctrl *AttrController) GetAttr(c *gin.Context) {
	// Convert the app ID to an integer
	appIdIntger, ok := GetAppID(c)
	if !ok {
		return
	}

	attrIdIntger, ok := GetAttrID(c)
	if !ok {
		return
	}

	// Query the app by ID
	attrDB := ctrl.Query.Attr
	attr, err := attrDB.Where(attrDB.ID.Eq(uint(attrIdIntger))).Where(attrDB.AppID.Eq(int32(appIdIntger))).First()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attr": attr,
	})
}

// @Summary 通用Attr创建
// @Description 用户通用属性创建，如果是文字和图片场景，请参考 对话Attr创建 和 图片Attr创建
// @Tags Attr
// @Accept json
// @Produce json
// @Param app_id path int true "App ID"
// @Param app body Attr true "Attr data"
// @Success 200 {object} Attr
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{app_id}/attrs [post]
func (ctrl *AttrController) CreateAttr(c *gin.Context) {
	// TODO 创建时需要判断 app_id是否存在
	appId, ok := GetAppID(c)
	if !ok {
		return
	}

	attr := model.Attr{}
	// Parse request body to extract app data
	err := c.BindJSON(&attr)
	if err != nil {
		common.Info(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	attr.AppID = int32(appId)

	// Validate app data
	if attr.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr name is required"})
		return
	}

	// Validate app data
	if attr.Type == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr type is required"})
		return
	}

	// Validate app data
	if attr.Context == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr context is required"})
		return
	}

	db := ctrl.Query.Attr
	// Save app data to database
	if err := db.Create(&attr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save app data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"attr": attr,
	})
}

// @Summary 对话Attr创建
// @Description 使用提供的数据创建新的 对话Attr
// @Tags Attr
// @Accept json
// @Produce json
// @Param app_id path int true "应用ID"
// @Param appData body ChatAttr true "更新的应用数据"
// @Success 200 {object} Attr
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{app_id}/chat_attrs [post]
func (ctrl *AttrController) CreateChatAttr(c *gin.Context) {
	// TODO 创建时需要判断 app_id是否存在
	appId, ok := GetAppID(c)
	if !ok {
		return
	}

	var chatAttr ChatAttr
	if err := c.BindJSON(&chatAttr); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	context, err := json.Marshal(chatAttr.Context)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	attr := model.Attr{
		Type:    chatAttr.Type,
		Name:    chatAttr.Name,
		Context: string(context),
	}

	attr.AppID = int32(appId)

	// Validate app data
	if attr.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr name is required"})
		return
	}

	// Validate app data
	if attr.Type == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr type is required"})
		return
	}

	// Validate app data
	if attr.Context == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Attr context is required"})
		return
	}

	db := ctrl.Query.Attr
	// Save app data to database
	if err := db.Create(&attr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save app data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attr": attr,
	})
}

// UpdateAttr godoc
// @Summary Update an attribute
// @Description Update an attribute by providing the app ID, attribute ID, and updated attribute information
// @Tags Attr
// @Accept json
// @Produce json
// @Param appID path int true "App ID" format(int32)
// @Param attrID path int true "Attribute ID" format(int32)
// @Param updatedAttr body Attr true "Updated attribute information"
// @Success 200 {object} common.SuccResponse "Success"
// @Failure 400 {object} common.ErrorResponse "Bad Request"
// @Failure 500 {object} common.ErrorResponse "Internal Server Error"
// @Router /v1/apps/{appID}/attrs/{attrID} [put]
func (ctrl *AttrController) UpdateAttr(c *gin.Context) {
	// Get the app ID from the URL parameter
	appID, ok := GetAppID(c)
	if !ok {
		return
	}
	attrId, ok := GetAttrID(c)
	if !ok {
		return
	}

	var updatedAttr *model.Attr
	err := c.ShouldBindJSON(&updatedAttr)
	if err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	db := ctrl.Query.Attr
	existingAttr, err := db.Where(db.ID.Eq(uint(attrId))).Where(db.AppID.Eq(int32(appID))).First()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	existingAttr.Name = updatedAttr.Name
	existingAttr.Context = updatedAttr.Context
	existingAttr.Type = updatedAttr.Type

	db.Save(existingAttr)

	c.JSON(http.StatusOK, gin.H{
		"app": existingAttr,
	})
}

// @Summary 对话Attr更新
// @Description 使用提供的数据创建新的 对话Attr
// @Tags Attr
// @Accept json
// @Produce json
// @Param app_id path int true "应用ID"
// @Param attr_id path int true "AttrID"
// @Param appData body ChatAttr true "更新的应用数据"
// @Success 200 {object} Attr
// @Failure 400 {object} common.ErrorResponse
// @Failure 500 {object} common.ErrorResponse
// @Router /v1/apps/{app_id}/chat_attrs/{attr_id} [put]
func (ctrl *AttrController) UpdateChatAttr(c *gin.Context) {
	// Get the app ID from the URL parameter
	appID, ok := GetAppID(c)
	if !ok {
		return
	}
	attrId, ok := GetAttrID(c)
	if !ok {
		return
	}

	db := ctrl.Query.Attr
	existingAttr, err := db.Where(db.ID.Eq(uint(attrId))).Where(db.AppID.Eq(int32(appID))).First()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	var chatAttr ChatAttr
	if err := c.BindJSON(&chatAttr); err != nil {
		common.Error(c, http.StatusBadRequest, err)
		return
	}

	context, err := json.Marshal(chatAttr.Context)
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	existingAttr.Name = chatAttr.Name
	existingAttr.Context = string(context)
	existingAttr.Type = chatAttr.Type

	// Save app data to database
	if err := db.Save(existingAttr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save app data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attr": existingAttr,
	})
}

// DeleteAttr godoc
// @Summary Delete an attribute
// @Description Delete an attribute by providing the app ID, attribute ID, and updated attribute information
// @Tags Attr
// @Accept json
// @Produce json
// @Param appID path int true "App ID" format(int32)
// @Param attrID path int true "Attribute ID" format(int32)
// @Success 200 {object} common.SuccResponse "Success"
// @Failure 400 {object} common.ErrorResponse "Bad Request"
// @Failure 500 {object} common.ErrorResponse "Internal Server Error"
// @Router /v1/apps/{appID}/attrs/{attrID} [delete]
func (ctrl *AttrController) DeleteAttr(c *gin.Context) {
	// Get the app ID from the URL parameter
	appID, ok := GetAppID(c)
	if !ok {
		return
	}
	attrId, ok := GetAttrID(c)
	if !ok {
		return
	}

	db := ctrl.Query.Attr
	existingAttr, err := db.Where(db.ID.Eq(uint(attrId))).Where(db.AppID.Eq(int32(appID))).First()
	if err != nil {
		common.Error(c, http.StatusInternalServerError, err)
		return
	}

	db.Delete(existingAttr)
	c.JSON(http.StatusOK, gin.H{
		"succ": true,
	})
}

func GetAppID(c *gin.Context) (int, bool) {
	appID := c.Param("app_id")

	if appID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing app ID",
		})
		return 0, false
	}
	// Convert the app ID to an integer
	id, err := strconv.Atoi(appID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid app id",
		})
		return 0, false
	}
	return id, true
}

func GetAttrID(c *gin.Context) (int, bool) {
	AttrID := c.Param("attr_id")

	if AttrID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "missing attr id",
		})
		return 0, false
	}
	// Convert the app ID to an integer
	id, err := strconv.Atoi(AttrID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid attr id",
		})
		return 0, false
	}
	return id, true
}

func init() {
	router := infra.GetApiEngine()
	attCtrl := NewAttrController()

	// Attr 相关路由
	router.GET("/apps/:app_id/attrs/", attCtrl.GetAttrs)
	router.GET("/apps/:app_id/attrs/:attr_id", attCtrl.GetAttr)
	router.POST("/apps/:app_id/attrs", attCtrl.CreateAttr)
	router.POST("/apps/:app_id/chat_attrs", attCtrl.CreateChatAttr)
	// router.POST("/apps/:app_id/img_attrs", attCtrl.CreateChatImagAttr)
	router.PUT("/apps/:app_id/attrs/:attr_id", attCtrl.UpdateAttr)
	router.PUT("/apps/:app_id/chat_attrs/:attr_id", attCtrl.UpdateChatAttr)
	router.DELETE("/apps/:app_id/attrs/:attr_id", attCtrl.DeleteAttr)

}
