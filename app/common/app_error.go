package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"error"`
	Msg        string `json:"msg"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func (t *AppError) Error() string {
	return t.RootErr.Error()
}
func NewSuccessResponse() *AppError {
	return &AppError{
		StatusCode: 200,
		RootErr:    nil,
		Msg:        "Success",
		Log:        "Success",
		Key:        "Success",
	}
}
func ErrorDB(err error) *AppError {
	return NewErrorResponse(err, "Something when wrong with database", err.Error(), "DB_ERROR")
}
func NewErrorResponse(err error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    err,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}
func NewFullErrorResponse(code int, err error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: code,
		RootErr:    err,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}
func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

// NewForbidden response status code 403
func NewForbidden(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusForbidden,
		RootErr:    root,
		Msg:        msg,
		Log:        msg,
		Key:        key,
	}
}

// NewUnauthorized response status code 401
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Msg:        msg,
		Log:        msg,
		Key:        key,
	}
}

func ErrorWrongConfirmPassword(err error) *AppError {
	msg := "Confirm password is incorrect. Please try again"
	key := "ErrWrongConfirmPassword"
	return NewCustomError(err, msg, key)
}

func NewCustomError(err error, msg, key string) *AppError {
	if err != nil {
		return NewErrorResponse(err, msg, msg, key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrWrongCurrentPassword(err error) *AppError {
	return NewCustomError(
		err,
		("Current password is incorrect. Please try again"),
		("ErrWrongPassword"),
	)
}

// ErrRecordNotFound is error message when gorm.ErrRecordNotFound called
var ErrRecordNotFound = errors.New("record not found")

// ErrEntityNotFound entity not found
func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}
