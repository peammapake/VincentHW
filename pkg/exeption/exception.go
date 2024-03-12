package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ExceptionError struct {
	ErrItem struct {
		Code    Code                   `json:"code"`
		Message string                 `json:"message"`
		Details map[string]interface{} `json:"details"`
	} `json:"error"`
}

func NewError(errCode Code, errMessage string) *ExceptionError {
	e := ExceptionError{}

	e.ErrItem.Code = errCode
	e.ErrItem.Message = errMessage

	return &e
}

func NewErrorWithDetails(errCode Code, errMessage string, errDetails map[string]string) *ExceptionError {
	e := ExceptionError{}

	e.ErrItem.Code = errCode
	e.ErrItem.Message = errMessage
	e.ErrItem.Details = make(map[string]interface{})

	for k, v := range errDetails {
		e.ErrItem.Details[k] = v
	}

	return &e
}

func NewValidationError(err error) *ExceptionError {
	e := ExceptionError{}

	e.ErrItem.Code = BAD_REQUEST
	e.ErrItem.Message = "Request validation error"
	e.ErrItem.Details = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		e.ErrItem.Details[v.Field()] = fmt.Sprintf("%v", v.Tag())
	}
	return &e
}
