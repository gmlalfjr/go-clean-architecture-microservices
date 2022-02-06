package exception

import "fmt"

type ErrorResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data string `json:"data"`
}


func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("status %d: err %v", r.Code, r.Data)
}

func BadRequestError(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:   400,
		Status: "Bad Request Error",
		Data:   err.Error(),
	}
}

func NewInternalServerError(err error) *ErrorResponse {
	return &ErrorResponse{
		Code:   500,
		Status: "Internal Server Error",
		Data:   err.Error(),
	}
}

func NewError(err error, code int, status string) *ErrorResponse {
	return &ErrorResponse{
		Code:   code,
		Status: status,
		Data:   err.Error(),
	}
}

