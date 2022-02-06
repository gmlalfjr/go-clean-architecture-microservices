package web

import "net/http"

type Response struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

func SuccessCreate(data interface{}) *Response {
	return &Response{
		Code:   http.StatusCreated,
		Status: "Successfully Create Data",
		Data:   data,
	}
}

