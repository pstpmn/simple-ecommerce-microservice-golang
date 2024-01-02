package pkg

import (
	"github.com/labstack/echo/v4"
)

type (
	IRespnse interface {
		Error(statusCode int, message string, ctx echo.Context) error
		Success(statusCode int, message string, result any, ctx echo.Context) error
	}

	Response struct {
		Message    string `json:"message"`
		Result     any    `json:"result,omitempty"`
		StatusCode int    `json:"statusCode"`
	}
)

type response struct {
}

// Error implements IRespnse.
func (*response) Error(statusCode int, message string, ctx echo.Context) error {
	return ctx.JSON(statusCode, &Response{
		Message:    message,
		StatusCode: statusCode,
	})
}

// Success implements IRespnse.
func (*response) Success(statusCode int, message string, result any, ctx echo.Context) error {
	return ctx.JSON(statusCode, &Response{
		Message:    message,
		Result:     result,
		StatusCode: statusCode,
	})
}

func NewResponse() IRespnse {
	return &response{}
}
