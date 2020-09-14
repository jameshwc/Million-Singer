package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
)

// ValidateUser godoc
// @Summary Check if username or email conflicts with one in database when registering
// @Description Check if username or email conflicts with one in database when registering
// @Tags user,check
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
	if username == "" {
		email := c.Query("email")
		if model.CheckDuplicateUserWithEmail(email) {
			appG.Response(http.StatusConflict, C.ERROR_REGISTER_EMAIL_CONFLICT, nil)
			return
		}
		appG.Response(http.StatusOK, C.SUCCESS, nil)
	} else {
		if model.CheckDuplicateUserWithName(username) {
			appG.Response(http.StatusConflict, C.ERROR_REGISTER_USERNAME_CONFLICT, nil)
			return
		}
		appG.Response(http.StatusOK, C.SUCCESS, nil)
	}
	appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, nil)
}
