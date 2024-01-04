package domain

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

var (
	ErrInternalServerError = errors.New("internal Server Error")
	ErrNotFound            = errors.New("your requested data is not found")
	ErrConflict            = errors.New("your data already exist")
	ErrBadParamInput       = errors.New("given Param is not valid")
	ErrBadRequest          = errors.New("given Request is not valid")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	case ErrBadRequest, ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func GenerateResponseError(err error, msg string, data interface{}) (res Response) {
	res.Code = GetStatusCode(err)
	if err == nil {
		res.Message = "Successfull"
		return res
	}

	if msg == "" {
		res.Message = err.Error()
	} else {
		res.Message = msg
	}
	res.Data = data

	return
}

func GenerateReponseSuccess(data interface{}) Response {
	return Response{
		Code:    0,
		Message: "successfull",
		Data:    data,
	}
}
