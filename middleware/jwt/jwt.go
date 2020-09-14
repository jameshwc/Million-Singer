package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/token"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}

		code := C.SUCCESS
		userToken := c.GetHeader("token")
		if userToken == "" {
			code = C.UNAUTHORIZED
		} else {
			_, err := token.Parse(userToken)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = C.ERROR_AUTH_TOKEN_TIMEOUT
				default:
					code = C.ERROR_AUTH_TOKEN_FAIL
				}
			}
		}

		if code != C.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  C.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
