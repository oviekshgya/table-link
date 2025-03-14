package helper

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

func NewErrorAuthLoginUnauthorized() AppError {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: "Username or password incorrect",
	}
}

func NewErrorUserNotFound() AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: "User not found",
	}
}

func NewErrorUserUsernameExist() AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "Username already exist",
	}
}

func NewErrorUserPasswordIncorrect() AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "Password incorrect",
	}
}

func NewErrorEmailExist() AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: "Email already exist",
	}
}
