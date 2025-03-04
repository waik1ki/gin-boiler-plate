package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: http.StatusText(code),
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(code, Response{
		Code:    code,
		Message: http.StatusText(code),
	})
}
