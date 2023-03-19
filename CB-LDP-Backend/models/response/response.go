package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"response_data,omitempty"`
}

func APIResponse(c *gin.Context, data interface{}, statusCode int, msg string) {
	c.JSON(statusCode, Response{Status: statusCode, Message: msg, Data: data})
}
