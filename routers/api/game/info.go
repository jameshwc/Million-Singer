package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service"
)

func GetSupportedLanguages(c *gin.Context) {
	appG := app.Gin{C: c}

	languages := service.Game.GetSupportedLanguages()

	appG.Response(http.StatusOK, constant.SUCCESS, constant.SuccessMsg, languages)
}

func GetGenres(c *gin.Context) {
	appG := app.Gin{C: c}

	genres := service.Game.GetGenres()

	appG.Response(http.StatusOK, constant.SUCCESS, constant.SuccessMsg, genres)
}
