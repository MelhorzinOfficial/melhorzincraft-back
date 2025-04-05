package response

import "github.com/MelhorzinOfficial/melhorzincraft-back/internal/oserror"

type Response[T any] struct {
	Code    int    `json:"-"`
	Error   bool   `json:"-"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data"`
}

func NewBadRequest[T any](message string) *Response[T] {
	return &Response[T]{
		Code:    400,
		Error:   true,
		Message: message,
	}
}

func NewNotFound[T any]() *Response[T] {
	return &Response[T]{
		Code:    404,
		Error:   true,
		Message: oserror.ErrNotFound.Error(),
	}
}

func NewConflict[T any](message string) *Response[T] {
	return &Response[T]{
		Code:    409,
		Error:   true,
		Message: message,
	}
}

func NewInternalServerError[T any]() *Response[T] {
	return &Response[T]{
		Code:    500,
		Error:   true,
		Message: oserror.ErrServer.Error(),
	}
}

func NewSuccess[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    200,
		Error:   false,
		Message: "success",
		Data:    data,
	}
}
