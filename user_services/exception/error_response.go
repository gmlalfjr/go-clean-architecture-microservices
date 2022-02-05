package exception

import "fmt"

type ErrorResponse struct {
	Code int
	Status string
	Data interface{}
}


func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("status %d: err %v", r.Code, r.Data)
}

func BadRequestError(data string) *ErrorResponse {
	return &ErrorResponse{
		Code:   400,
		Status: "Bad Request Error",
		Data:   data,
	}
}

func NewInternalServerError(data string) *ErrorResponse {
	return &ErrorResponse{
		Code:   500,
		Status: "Internal Server Error",
		Data:   data,
	}
}
