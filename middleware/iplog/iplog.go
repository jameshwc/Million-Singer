package iplog

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/log"
)

func TraceIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.TraceIP(c.ClientIP(), c.Request.RequestURI)
		c.Next()
	}
}
