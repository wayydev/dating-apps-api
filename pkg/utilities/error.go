package utilities

import "github.com/gin-gonic/gin"

type Error struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Stack   interface{} `json:"stack"`
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			var err *Error
			if r := recover(); r != nil {
				switch t := r.(type) {
				case string:
					err = &Error{Message: t, Status: 500}
				case *Error:
					err = t
				case error:
					err = &Error{Message: t.Error(), Status: 500}
				default:
					err = &Error{Message: "Unknow Error", Status: 500}
				}
				c.JSON(err.Status, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func ErrorNotFound(c *gin.Context) {
	c.JSON(404, &Error{
		Message: "Page Not Found",
		Status:  404,
		Data:    nil,
		Stack:   nil,
	})
}
