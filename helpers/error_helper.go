package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"ralali.com/constants"
	"syscall"
	"time"
)

type ErrorHelper struct {
}

func ErrorHelperHandler() (ErrorHelper) {
	return ErrorHelper{}
}

func (handler *ErrorHelper) HTTPResponseError(context *gin.Context, e error, errorCode int) {
	errorConstant := constants.GetErrorConstant(errorCode)
	context.JSON(errorConstant.HttpCode, gin.H{
		"code":    errorConstant.HttpCode,
		"message": errorConstant.Message,
	})
	handler.LogError(e, true)
}

func (handler *ErrorHelper) LogError(e error, isPanic bool) {

	currentTime := time.Now()

	filePath := os.Getenv("ERROR_LOG_FILE")
	fileName := currentTime.Format("2006-01-02")

	logFullPath := fmt.Sprintf("%s/error-%s.log", filePath, fileName)

	file, err := os.OpenFile(logFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if nil != err {
		fmt.Println(err)
	}
	defer file.Close()

	syscall.Dup2(int(file.Fd()), int(os.Stderr.Fd()))

	if nil == e {
		log.Println("There's no error but calling logError!")
	} else {
		log.Println(e)
	}

	if isPanic {
		log.Panic(e.Error() + "\n")
	}

}
