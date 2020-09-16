package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	gameService "github.com/jameshwc/Million-Singer/service/game"
)

type level struct {
	Songs []int  `json:"songs"`
	Title string `json:"title"`
}

// AddLevel godoc
// @Summary Add a new level
// @Description Add a new level
// @Tags game,level
// @Accept json
// @Produce json
// @Param token header string true "auth token, must register & login to get the token"
// @Param songs json string true "id of the song, should have many"
// @Param title json string true "title of the level"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/levels/new [post]
func AddLevel(c *gin.Context) {

	appG := app.Gin{C: c}
	var l level

	if err := c.BindJSON(&l); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, nil)
		return
	}

	switch err := gameService.AddLevel(l.Songs, l.Title); err {

	case C.ErrLevelAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_NO_SONGID_OR_TITLE, nil)

	case C.ErrLevelAddSongsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_NO_SONGID_RECORD, nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_ADD_LEVEL_SERVER_ERROR, nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, nil)

	}
}

// GetLevel godoc
// @Summary Get a level
// @Description Get a level
// @Tags game,level
// @Accept plain
// @Produce json
// @Param id path int true "id of the level"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/levels/{id} [get]
func GetLevel(c *gin.Context) {

	appG := app.Gin{C: c}

	switch level, err := gameService.GetLevel(c.Param("id")); err {

	case C.ErrLevelIDNotNumber:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_LEVEL_ID_NAN, nil)

	case C.ErrLevelNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_LEVEL_NO_RECORD, nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, level)

	}
}
