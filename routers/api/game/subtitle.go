package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service"
)

type subtitleFile struct {
	Filetype string `json:"file_type"`
	File     []byte `json:"file"`
}

func ConvertFileToSubtitle(c *gin.Context) {
	appG := app.Gin{C: c}

	var param subtitleFile
	if err := c.BindJSON(&param); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)
		return
	}

	switch sub, err := service.Game.ConvertFileToSubtitle(param.Filetype, param.File); err {

	case C.ErrConvertFileToSubtitleTypeNotSupported:
		appG.Response(http.StatusBadRequest, C.ERROR_CONVERT_FILE_TYPE_NOT_SUPPORTED, err.Error(), nil)

	case C.ErrConvertFileToSubtiteParse:
		appG.Response(http.StatusBadRequest, C.ERROR_CONVERT_FILE_PARSE, err.Error(), nil)

	case nil:
		appG.Response(http.StatusAccepted, C.SUCCESS, C.SuccessMsg, sub)

	}
}

func DownloadYoutubeSubtitle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch sub, err := service.Game.DownloadYoutubeSubtitle(c.Query("url"), c.Query("code")); err {

	case C.ErrDownloadYoutubeSubtitle:
		appG.Response(http.StatusBadRequest, C.ERROR_DOWNLOAD_YOUTUBE_SUBTITLE, err.Error(), nil)

	case nil:
		appG.Response(http.StatusAccepted, C.SUCCESS, C.SuccessMsg, sub)

	}
}

func GetYoutubeTitle(c *gin.Context) {

	appG := app.Gin{C: c}

	switch title, err := service.Game.GetYoutubeTitle(c.Query("url")); err {

	case C.ErrGetYoutubeTitle:
		appG.Response(http.StatusBadRequest, C.ERROR_DOWNLOAD_YOUTUBE_SUBTITLE, err.Error(), nil)

	case nil:
		appG.Response(http.StatusAccepted, C.SUCCESS, C.SuccessMsg, title)

	}
}
