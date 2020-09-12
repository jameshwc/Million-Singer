package app

import (
	"github.com/gin-gonic/gin"

	"github.com/jameshwc/Million-Singer/pkg/constant"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"success"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  constant.GetMsg(errCode),
		Data: data,
	})
	return
}
