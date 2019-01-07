package controllers

import (
	// "../constants"
	"../helpers"
	"../objects"
	"../services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type V2ShortyController struct {
	shortyService services.V2ShortyService
	errorHelper helpers.ErrorHelper
}

func V2ShortyControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &V2ShortyController{
		shortyService: services.V2ShortyServiceHandler(db),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("/")
	{
		group.POST("shorten", handler.PostByShorten)
	}

}

func (handler *V2ShortyController) PostByShorten(context *gin.Context) {
	requestObject := objects.V2ShortyObjectRequest{}
	context.ShouldBind(&requestObject)

	validateShortcode := handler.shortyService.ValidateShortcode(requestObject.Shortcode)
	if validateShortcode != true {
		context.JSON(422, gin.H{
			"code": 422,
			"message": "The shortcode must contain exactly 6 alpahnumeric characters",
		})
		return
	}
	
	validateUrl := handler.shortyService.ValidateUrl(requestObject.Url)
	if validateUrl != true {
		context.JSON(400, gin.H{
			"code": 400,
			"message": "Url is not present",
		})
		return
	}

	
	result, err := handler.shortyService.PostByShorten(requestObject)
	if nil != err {
		context.JSON(409, gin.H{
			"code": 409,
			"message": "The desired shortcode is already in use",
		})
		return
	}

	context.JSON(http.StatusOK, result)
}
