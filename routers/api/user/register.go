package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	userService "github.com/jameshwc/Million-Singer/service/user"
)

// Register godoc
// @Summary Register a user
// @Description Register a user
// @Tags user
// @Accept json
// @Produce json
// @Param username body string true "username"
// @Param email body string true "email"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/register [get]
func Register(c *gin.Context) {
	appG := app.Gin{C: c}

	var u user
	c.BindJSON(&u)

	switch err := userService.Register(u.Username, u.Email, u.Password); err {

	case C.ErrUserRegisterFormat:
		appG.Response(http.StatusBadRequest, C.ERROR_REGISTER_FORMAT_INCORRECT, nil)

	case C.ErrUserRegisterNameConflict:
		appG.Response(http.StatusConflict, C.ERROR_REGISTER_USERNAME_CONFLICT, nil)

	case C.ErrUserRegisterEmailConflict:
		appG.Response(http.StatusConflict, C.ERROR_REGISTER_EMAIL_CONFLICT, nil)

	case C.ErrUserRegisterFailServerError:
		appG.Response(http.StatusInternalServerError, C.ERROR_REGISTER_FAIL_SERVER_ERROR, nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, nil)
	}
}
