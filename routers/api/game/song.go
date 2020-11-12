package game

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service"
	gameService "github.com/jameshwc/Million-Singer/service/game"
)

// AddSong godoc
// @Summary Add a song
// @Description Add a song to database
// @Tags game,song
// @Accept json
// @Produce json
// @Param token header string true "auth token, must register & login to get the token"
// @Param file body string true "subtitle file"
// @Param url body string true "youtube url"
// @Param name body string true "name of the song"
// @Param singer body string true "singer of the song"
// @Param miss_lyrics body string true "miss lyrics index, seperated by commas"
// @Param genre body string false "genre of the song"
// @Param language body string false "language of the song, needs to be short, and seperated by commas if multiple languages"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/songs/new [post]
func AddSong(c *gin.Context) {
	appG := app.Gin{C: c}

	var song *gameService.Song

	c.BindJSON(&song)

	switch songID, err := service.Game.AddSong(song); err {

	case C.ErrSongFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_SONG_FORMAT_INCORRECT, err.Error(), nil)

	case C.ErrSongAddLyricsFileTypeNotSupported:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_SONG_LYRICS_FILE_TYPE_NOT_SUPPORTED, err.Error(), nil)

	case C.ErrSongAddParseLyrics:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_SONG_PARSE_LYRICS_ERROR, err.Error(), nil)

	case C.ErrSongAddMissLyricsIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_SONG_PARSE_LYRICS_ERROR, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_ADD_SONG_SERVER_ERROR, err.Error(), nil)

	case C.ErrSongAddDuplicate:
		appG.Response(http.StatusUnprocessableEntity, C.ERROR_ADD_SONG_DUPLICATE, fmt.Errorf("%w, song id: %d", err, songID).Error(), nil)

	case C.ErrSongAddURLIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_SONG_URL_INCORRECT, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, songID)

	}

}

// GetSong godoc
// @Summary Get a song Instance
// @Description Get a song instance, with a miss lyrics id randomly generated
// @Tags game,song
// @Accept json
// @Produce json
// @Param id path int true "id of the song"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/songs/{id} [get]
func GetSongInstance(c *gin.Context) {
	appG := app.Gin{C: c}

	var hasLyrics bool
	if queryLyrics := c.DefaultQuery("lyrics", "y"); queryLyrics == "y" {
		hasLyrics = true
	}

	switch s, err := service.Game.GetSongInstance(c.Param("id"), hasLyrics); err {

	case C.ErrSongIDNotNumber:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_SONG_ID_NAN, err.Error(), nil)

	case C.ErrSongNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_SONG_NO_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_SONG_SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, s)

	}
}

func DeleteSong(c *gin.Context) {
	appG := app.Gin{C: c}
	switch err := service.Game.DeleteSong(c.Param("id")); err {

	case C.ErrSongIDNotNumber:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_SONG_ID_NAN, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_SONG_SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)

	}
}

func GetSongs(c *gin.Context) {
	appG := app.Gin{C: c}
	switch songs, err := service.Game.GetSongs(); err {
	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_SONG_SERVER_ERROR, err.Error(), nil)
	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, songs)
	}
}
