package controllers

import (
	"../constants"
	"../helpers"
	"../models"
	"../repositories"
	"../requests"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	repository    repositories.UserRepository
	errorHandling helpers.ErrorHandling
	request       requests.UserRequest

	orderRepository repositories.OrderRepository
}

func UserControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &UserController{
		repository:      repositories.UserRepository{DB: *db},
		errorHandling:   helpers.ErrorHandling{},
		request:         requests.UserRequest{},
		orderRepository: repositories.OrderRepository{DB: *db},
	}

	group := router.Group("users")
	{
		group.GET("", handler.GetList)
		group.GET(":id", handler.GetById)
		group.GET(":id/orders", handler.GetOrderByUserId)
		group.POST(":id", handler.UpdateUser)
	}

}

func (handler *UserController) GetOrderByUserId(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	request := handler.request.GetList
	request.Page = 1
	request.PerPage = 50

	err = context.ShouldBind(&request)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	result, err := handler.orderRepository.GetListByUserId(id, request.Page, request.PerPage)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, result)

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