package utils

import (
	"errors"
	"net/http"
	"strings"
)

//Response is used for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(data interface{}) Response {
	res := Response{
		Status:  true,
		Message: "API Success",
		Errors:  nil,
		Data:    data,
	}
	return res
}

type CartResponse struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
	Data       interface{} `json:"data"`
	TotalPrice float64     `json:"total_price"`
}

//BuildCartResponse method is to inject data value to list cart response and total price
func BuildCartResponse(data interface{}, totalPrice float64) CartResponse {
	res := CartResponse{
		Status:     true,
		Message:    "API Success",
		Errors:     nil,
		Data:       data,
		TotalPrice: totalPrice,
	}
	return res
}

func BuildSignOutResponse() Response {
	res := Response{
		Status:  true,
		Message: "Logout Successful",
		Errors:  nil,
		Data:    EmptyObj{},
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(err interface{}) Response {
	//splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: "API Failed",
		Errors:  err,
		Data:    EmptyObj{},
	}
	return res
}

func BuildValidationErrorResponse() Response {
	splittedError := strings.Split("Validation Error", "\n")
	res := Response{
		Status:  false,
		Message: "API Failed",
		Errors:  splittedError,
		Data:    EmptyObj{},
	}
	return res
}

func BuildAuthErrorResponse() Response {
	res := Response{
		Status:  false,
		Message: "API Failed",
		Errors:  NewStatusUnAuthorizedError(),
		Data:    EmptyObj{},
	}
	return res
}

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

func NewStatusUnAuthorizedError() *RestErr {
	return &RestErr{
		Message: "JWT token authentication failed",
		Status:  http.StatusUnauthorized,
		Error:   "Status UnAuthorized",
	}
}
