package utils

import "bookstore/internal/app/model"

func Success(data any, message string) model.Response {
	return model.Response{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func Created(data any, message string) model.Response {
	return model.Response{
		Code:    201,
		Message: message,
		Data:    data,
	}
}

func Error(message string, code int) model.Response {
	return model.Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}