package controllers

import (
	// "../constants"
	"../helpers"
	// "../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type V1ShortyController struct {
	shortyService services.V1ShortyService
	errorHelper helpers.ErrorHelper
}

func V1ShortyControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &V1ShortyController{
		shortyService: services.V1ShortyServiceHandler(db),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("/")
	{
		group.GET(":shortcode", handler.GetByShortcode)
		group.GET(":shortcode/stats", handler.GetByShortcodeStats)
	}

}

func (handler *V1ShortyController) GetByShortcode(context *gin.Context) {

	var shortcode = context.Param("shortcode")
	result, err := handler.shortyService.GetByShortcode(shortcode)
	if nil != err {
		context.JSON(404, gin.H{
			"code": 404,
			"message": "The shortcode cannot be found in the system",
		})
		return
	}
	context.Redirect(301,result.Url)

}

func (handler *V1ShortyController) GetByShortcodeStats(context *gin.Context) {

	var shortcode = context.Param("shortcode")
	result, err := handler.shortyService.GetByShortcodeStats(shortcode)
	if nil != err {
		context.JSON(404, gin.H{
			"code": 404,
			"message": "The shortcode cannot be found in the system",
		})
		return
	}
	context.JSON(http.StatusOK, result)
}
