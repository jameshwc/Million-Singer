package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
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
		appG.Response(http.StatusBadRequest, constant.ERROR_ADD_LEVEL_NO_SONGID_PARAM, nil)
		return
	}
	for i := range songsIDstr {
		songID, err := strconv.Atoi(songsIDstr[i])
		if err != nil {
			appG.Response(http.StatusBadRequest, constant.ERROR_ADD_LEVEL_SONG_NAN, nil)
			return
		}
		songsID = append(songsID, songID)
	}
	// level.SongsID = songsID
	level.Title, check = c.GetPostForm("title")
	if !check {
		appG.Response(http.StatusBadRequest, constant.ERROR_ADD_LEVEL_NO_TITLE, nil)
		return
	}
	var err error
	if level.Songs, err = model.GetSongs(songsID); err != nil {
		appG.Response(http.StatusBadRequest, constant.ERROR_ADD_LEVEL_NO_SONGID_RECORD, nil)
		return
	}
	if err = level.Commit(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_ADD_LEVEL_SERVER_ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, level)
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
// @Router /game/levels/{id} [post]
func GetLevel(c *gin.Context) {
	appG := app.Gin{C: c}
	levelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.ERROR_GET_LEVEL_FAIL, nil)
		return
	}
	level, err := model.GetLevel(levelID)
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.ERROR_GET_LEVEL_NO_RECORD, nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, level)
}
