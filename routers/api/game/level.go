package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	gameService "github.com/jameshwc/Million-Singer/service/game"
)

// AddLevel godoc
// @Summary Add a new level
// @Description Add a new level
// @Tags game,level
// @Accept multipart/form-data
// @Produce json
// @Param token header string true "auth token, must register & login to get the token"
// @Param songs formData string true "id of the song, should have many"
// @Param title formData string true "title of the level"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/levels/new [post]
func AddLevel(c *gin.Context) {
	appG := app.Gin{C: c}
	var level model.Level
	var songsID []int
	songsIDstr, check := c.GetPostFormArray("songs")
	if !check {
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_NO_SONGID_PARAM, nil)
		return
	}
	for i := range songsIDstr {
		songID, err := strconv.Atoi(songsIDstr[i])
		if err != nil {
			appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_SONG_NAN, nil)
			return
		}
		songsID = append(songsID, songID)
	}
	// level.SongsID = songsID
	level.Title, check = c.GetPostForm("title")
	if !check {
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_NO_TITLE, nil)
		return
	}
	var err error
	if level.Songs, err = model.GetSongs(songsID); err != nil {
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_LEVEL_NO_SONGID_RECORD, nil)
		return
	}
	if err = level.Commit(); err != nil {
		appG.Response(http.StatusInternalServerError, C.ERROR_ADD_LEVEL_SERVER_ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, C.SUCCESS, level)
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
