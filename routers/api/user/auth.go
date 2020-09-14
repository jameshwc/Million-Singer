package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
)

// Login godoc
// @Summary Log in a user
// @Description Log in a user
// @Tags user,login
// @Accept plain
// @Produce json
// @Param username path string true "username"
// @Param password path string true "password"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, nil)
		return
	}
	u, err := model.AuthUser(username, password)
	if err != nil {
		appG.Response(http.StatusUnauthorized, C.ERROR_LOGIN_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, C.SUCCESS, u.Token)
}
