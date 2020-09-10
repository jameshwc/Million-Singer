package game

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/app"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
)

type SongInstance struct {
	Song        *model.Song
	MissLyricID int `json:"miss_lyric_id"`
}

func AddSong(c *gin.Context) {
	appG := app.Gin{C: c}
	var song model.Song
	// TODO: should seperate the model from the routers?
	song.URL = c.PostForm("url")
	song.Genre = c.PostForm("genre")
	song.Language = c.PostForm("language")
	song.MissLyrics = c.PostForm("miss_lyrics")
	song.Singer = c.PostForm("singer")
	song.Name = c.PostForm("name")
	srtFile, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		appG.Response(constant.INVALID_PARAMS, constant.ERROR_UPLOAD_SRT_FILE, nil)
		return
	}
	if err != nil {
		appG.Response(constant.SERVER_ERROR, constant.ERROR_UPLOAD_SRT_FILE, nil)
		return
	}
	song.Lyrics, err = subtitle.ReadSrtFromFile(srtFile)
	if err != nil {
		appG.Response(constant.INVALID_PARAMS, constant.ERROR_SRT_FILE_FORMAT, nil)
		return
	}
	if err := song.Commit(); err != nil {
		appG.Response(constant.SERVER_ERROR, constant.ERROR_ADD_SONG_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, constant.SUCCESS, song.ID)
}

func GetSongInstance(c *gin.Context) {
	appG := app.Gin{C: c}
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, constant.INVALID_PARAMS, nil)
		return
	}
	var hasLyrics bool
	if queryLyrics := c.DefaultQuery("lyrics", "y"); queryLyrics == "y" {
		hasLyrics = true
	}
	if s, err := model.GetSong(songID, hasLyrics); err != nil {
		appG.Response(http.StatusInternalServerError, constant.ERROR_GET_SONG_FAIL, nil)
	} else {
		var songInstance SongInstance
		songInstance.Song = s
		songInstance.MissLyricID = s.RandomGetMissLyricID()
		appG.Response(http.StatusOK, constant.SUCCESS, songInstance)
	}
}
