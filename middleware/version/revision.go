package version

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/conf"
)

func RevisionMiddleware() gin.HandlerFunc {
	revision := conf.ServerConfig.Revision
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Revision", revision)
		c.Next()
	}
}
