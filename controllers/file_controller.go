package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"ralali.com/constants"
	"ralali.com/helpers"
	"ralali.com/requests"
	"ralali.com/services"
)

type FileController struct {
	errorHandling helpers.ErrorHandling
	awsService    services.AWSS3Service
	request       requests.FileRequest
}

func FileControllerHandler(router *gin.Engine, db *gorm.DB) {

	handler := &FileController{
		errorHandling: helpers.ErrorHandling{},
		awsService:    services.AWSS3Service{},
		request:       requests.FileRequest{},
	}

	group := router.Group("files")
	{
		group.POST("", handler.UploadFile)
	}

}

func (handler *FileController) UploadFile(context *gin.Context) {

	request := handler.request.UploadFile
	request.SavePath = os.Getenv("S3_DEFAULT_PATH")

	err := context.ShouldBind(&request)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	}

	fileRequest, err := context.FormFile("file")
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.RequestParameterInvalid)
	} else {
		request.File = fileRequest
	}

	if "" == request.FileName {
		request.FileName = fileRequest.Filename
	}

	fileData, err := request.File.Open()
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	handler.awsService, _ = handler.awsService.Initialize()
	result, err := handler.awsService.UploadFile(fileData, request.SavePath, request.FileName)
	if nil != err {
		handler.errorHandling.HTTPResponseError(context, err, constants.InternalServerError)
	}

	context.JSON(http.StatusOK, gin.H{
		"s3_tag": result.ETag,
	})

}
