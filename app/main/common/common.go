package common

import (
	"net/http"
	"strings"
)

const (
	EmptyObject ErrorDataType = iota
	EmptyList
	EmptyString
	Zero
)

type (
	ErrorDataType int

	Response struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func NewResponse(msg string, data interface{}) (res Response) {
	res.Message = msg
	res.Data = data
	return
}

func getDataType(dataType ErrorDataType) interface{} {
	switch dataType {
	case EmptyList:
		return []string{}
	case EmptyString:
		return ""
	case Zero:
		return 0
	default:
		return map[string]interface{}{}
	}
}

func NewBadRequestResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusBadRequest)
	res.Data = getDataType(dataType)
	return
}

func NewUnauthorizedResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusUnauthorized)
	res.Data = getDataType(dataType)
	return
}

func NewInternalServerErrorResponse(dataType ErrorDataType) (res Response) {
	res.Message = http.StatusText(http.StatusInternalServerError)
	res.Data = getDataType(dataType)
	return
}

func NewNotFoundResponse(dataType ErrorDataType, parameters []string) (res Response) {
	res.Message = strings.Join(parameters, ", ") + " not found"
	res.Data = getDataType(dataType)
	return
}
