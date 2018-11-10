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

type UserController struct {
	userService services.UserService
	errorHelper helpers.ErrorHelper
}

func UserControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &UserController{
		userService: services.UserServiceHandler(db),
		errorHelper: helpers.ErrorHelperHandler(),
	}

	group := router.Group("users")
	{
		group.GET(":id", handler.GetById)
		group.POST(":id", handler.UpdateById)
	}

}

func (handler *UserController) GetById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.userService.GetById(id)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}

func (handler *UserController) UpdateById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	requestObject := objects.UserObject{}
	context.ShouldBind(&requestObject)

	result, err := handler.userService.UpdateById(id, requestObject)
	if nil != err {
		handler.errorHelper.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
