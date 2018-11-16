package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ralali.com/constants"
	"ralali.com/helpers"
	"ralali.com/objects"
	"ralali.com/services"
	"strconv"
)

type V2UserController struct {
	userService services.V2UserService
	errorHelper helpers.ErrorHelper
}

func V2UserControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &V2UserController{
		userService: services.V2UserServiceHandler(db),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("v2/users")
	{
		group.POST(":id", handler.UpdateById)
	}

}

func (handler *V2UserController) UpdateById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	requestObject := objects.V2UserObjectRequest{}
	context.ShouldBind(&requestObject)

	result, err := handler.userService.UpdateById(id, requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
