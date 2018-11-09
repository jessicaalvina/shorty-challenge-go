package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"ralali.com/constants"
	"ralali.com/helpers"
	"ralali.com/models"
	"ralali.com/repositories"
	"ralali.com/requests"
	"strconv"
)

type UserController struct {
	repository    repositories.UserRepository
	errorHandling helpers.ErrorHandling
	request       requests.UserRequest
}

func UserControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &UserController{
		repository:    repositories.UserRepository{DB: *db},
		errorHandling: helpers.ErrorHandling{},
		request:       requests.UserRequest{},
	}

	group := router.Group("users")
	{
		group.GET("", handler.GetList)
		group.GET(":id", handler.GetById)
		group.POST(":id", handler.UpdateUser)
	}

}

func (handler *UserController) GetById(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.repository.GetById(id)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}

func (handler *UserController) UpdateUser(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	user := models.User{}

	err = context.ShouldBindJSON(&user)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.repository.UpdateUser(id, user)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}

func (handler *UserController) GetList(context *gin.Context) {

	request := handler.request.GetList
	request.Page = 1
	request.PerPage = 50

	err := context.ShouldBind(&request)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.repository.GetList(request.Page, request.PerPage)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

}
