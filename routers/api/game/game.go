package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
)

func GetGame(c *gin.Context) {
	appG := app.Gin{C: c}
	gameID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}
	if g, err := model.GetGame(gameID); err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_GET_GAME_FAIL, nil)
	} else {
		appG.Response(http.StatusOK, constant.SUCCESS, g)
	}
}
