package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service"
)

// ValidateUser godoc
// @Summary Check if username or email conflicts with one in database when registering
// @Description Check if username or email conflicts with one in database when registering
// @Tags user
// @Accept plain
// @Produce json
// @Param name path string false "username that needs to check"
// @Param email path string false "email that needs to check"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /users/check [get]
func ValidateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	username := c.Query("name")
	email := c.Query("email")

	switch err := service.User.Validate(username, email); err {

	case C.ErrUserCheckParamIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_CHECK_PARAM_INCORRECT, err.Error(), nil)

	case C.ErrUserCheckFormat:
		appG.Response(http.StatusBadRequest, C.ERROR_CHECK_FORMAT_INCORRECT, err.Error(), nil)

	case C.ErrUserCheckEmailConflict:
		appG.Response(http.StatusConflict, C.ERROR_CHECK_EMAIL_CONFLICT, err.Error(), nil)

	case C.ErrUserCheckNameConflict:
		appG.Response(http.StatusConflict, C.ERROR_CHECK_NAME_CONFLICT, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)
	}

}
