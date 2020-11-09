package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg,omitempty" example:"success"`
	Data interface{} `json:"data,omitempty"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
