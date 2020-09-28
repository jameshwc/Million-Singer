package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service/game"
)

// GetLyricsWithSongID godoc
// @Summary Get lyrics with a song's ID
// @Description Get lyrics with a song's ID; normally it is only for internal use.
// @Tags game,lyric
// @Accept json
// @Produce json
// @Param song_id path int true "id of the song"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/lyrics/{song_id} [get]
func GetLyricsWithSongID(c *gin.Context) {
	appG := app.Gin{C: c}

	switch lyrics, err := game.GetLyricsWithSongID(c.Param("song_id")); err {

	case constant.ErrSongIDNotNumber:
		appG.Response(http.StatusBadRequest, constant.ERROR_GET_SONG_ID_NAN, err.Error())

	case constant.ErrSongNotFound:
		appG.Response(http.StatusBadRequest, constant.ERROR_GET_SONG_NO_RECORD, err.Error())

	case nil:
		appG.Response(http.StatusOK, constant.SUCCESS, lyrics)

	}
}
