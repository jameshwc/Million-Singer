package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
)

func GetLyricsWithSongID(c *gin.Context) {
	appG := app.Gin{C: c}
	songID, err := strconv.Atoi(c.Query("song_id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}
	if _, err := model.GetSong(songID, false); err != nil {
		appG.Response(http.StatusBadRequest, constant.ERROR_GET_SONG_FAIL, nil)
		return
	}
	if lyrics, err := model.GetLyricsWithSongID(songID); err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_GET_SONG_FAIL, nil)
	} else {
		appG.Response(http.StatusOK, constant.SUCCESS, lyrics)
	}
}
