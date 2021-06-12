package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/service"
)

type collect struct {
	Songs []int  `json:"songs"`
	Title string `json:"title"`
}

// AddCollect godoc
// @Summary Add a new Collect
// @Description Add a new Collect
// @Tags game,Collect
// @Accept json
// @Produce json
// @Param token header string true "auth token, must register & login to get the token"
// @Param songs body string true "id of the song, should have many"
// @Param title body string true "title of the Collect"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/collects/new [post]
func AddCollect(c *gin.Context) {

	appG := app.Gin{C: c}
	var l collect

	if err := c.BindJSON(&l); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)
		return
	}

	switch id, err := service.Game.AddCollect(l.Songs, l.Title); err {

	case C.ErrCollectAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_COLLECT_NO_SONGID_OR_TITLE, err.Error(), nil)

	case C.ErrCollectAddSongsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_COLLECT_NO_SONGID_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_ADD_COLLECT_SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, id)

	}
}

// GetCollect godoc
// @Summary Get a Collect
// @Description Get a Collect
// @Tags game,Collect
// @Accept plain
// @Produce json
// @Param id path int true "id of the Collect"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/Collects/{id} [get]
func GetCollect(c *gin.Context) {

	appG := app.Gin{C: c}

	switch Collect, err := service.Game.GetCollect(c.Param("id")); err {

	case C.ErrCollectIDNotNumber:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_COLLECT_ID_NAN, err.Error(), nil)

	case C.ErrCollectNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_COLLECT_NO_RECORD, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, Collect)

	}
}

func GetCollects(c *gin.Context) {
	appG := app.Gin{C: c}

	switch collects, err := service.Game.GetCollects(); err {
	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_SONG_SERVER_ERROR, err.Error(), nil)
	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, collects)
	}
}

func DelCollect(c *gin.Context) {
	appG := app.Gin{C: c}

	switch toursID, err := service.Game.DelCollect(c.Param("id")); err {

	case C.ErrCollectDelIDIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_DEL_COLLECT_ID_INCORRECT, err.Error(), nil)

	case C.ErrCollectDelDeleted:
		appG.Response(http.StatusGone, C.ERROR_DEL_COLLECT_DELETED, err.Error(), nil)

	case C.ErrCollectDelForeignKey:
		appG.Response(http.StatusUnprocessableEntity, C.ERROR_DEL_COLLECT_FOREIGN_KEY, err.Error(), toursID)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)

	}
}
