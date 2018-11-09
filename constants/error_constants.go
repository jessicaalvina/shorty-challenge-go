package constants

import "net/http"

type ErrorConstant struct {
	HttpCode int
	Message  string
}

const (
	RequestParameterInvalid      = 1000
	ObjectNotInitializedProperly = 1001
	InternalServerError          = 1002
)

var errorConstantMapping = map[int]ErrorConstant{
	RequestParameterInvalid: {
		HttpCode: http.StatusBadRequest,
		Message:  "Invalid request parameter",
	},
	ObjectNotInitializedProperly: {
		HttpCode: http.StatusInternalServerError,
		Message:  "Object is not initialized properly",
	},
	InternalServerError: {
		HttpCode: http.StatusInternalServerError,
		Message:  "Something went wrong",
	},
}

func GetErrorConstant(code int) ErrorConstant {
	return errorConstantMapping[code]
}
