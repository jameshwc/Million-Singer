package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	userService "github.com/jameshwc/Million-Singer/service/user"
)

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user
// @Tags user
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
	var u user
	c.BindJSON(&u)
	token, err := userService.AuthUser(u.Username, u.Password)
	switch err {

	case C.ErrUserLoginFormat:
		appG.Response(http.StatusBadRequest, C.ERROR_LOGIN_FAIL_FORMAT_INCORRECT, nil)

	case C.ErrUserLoginAuthentication:
		appG.Response(http.StatusBadRequest, C.ERROR_LOGIN_FAIL_AUTHENTICATION, nil)

	case C.ErrUserLoginJwtTokenGeneration:
		appG.Response(http.StatusInternalServerError, C.ERROR_LOGIN_FAIL_JWT_TOKEN_GENERATION, nil)

	case C.ErrUserLoginUpdateUserStatus:
		appG.Response(http.StatusInternalServerError, C.ERROR_LOGIN_FAIL_UPDATE_LOGIN_STATUS, nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, token)
	}
}
