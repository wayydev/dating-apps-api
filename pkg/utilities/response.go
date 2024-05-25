package utilities

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseJSON(c *gin.Context, message string, status int, data interface{}) {
	c.JSON(status, &Response{
		Message: message,
		Status:  status,
		Data:    data,
	})
}
