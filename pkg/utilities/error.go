package utilities

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorReponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Stack   interface{} `json:"stack"`
}

func (e *ErrorReponse) Error() string {
	return fmt.Sprintf("status: %d, message: %s", e.Status, e.Message)
}

func Error(message string, status int, data interface{}, stack interface{}) *ErrorReponse {
	return &ErrorReponse{
		Message: message,
		Status:  status,
		Data:    data,
		Stack:   stack,
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var err *ErrorReponse
			if r := recover(); r != nil {
				switch t := r.(type) {
				case string:
					err = &ErrorReponse{Message: t, Status: 500}
				case *ErrorReponse:
					err = t
				case validator.ValidationErrors:
					errors := TranslatorValidate(t)
					err = &ErrorReponse{Message: "Validation Error", Status: 422, Data: errors}
				case error:
					err = &ErrorReponse{Message: t.Error(), Status: 500, Stack: t}
				default:
					err = &ErrorReponse{Message: "Unknow Error", Status: 500}
				}
				c.JSON(err.Status, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func ErrorNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			c.JSON(404, &ErrorReponse{
				Message: "Page Not Found",
				Status:  404,
				Data:    nil,
				Stack:   nil,
			})
		}()
	}
}
