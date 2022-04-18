package model

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func NewResponse[T any](status string, message string, data T) *Response[T] {
	return &Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
