package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	gameService "github.com/jameshwc/Million-Singer/service/game"
)

type tour struct {
	Collects []int `json:"collects"`
}

// GetTour godoc
// @Summary Get a tour by id
// @Description Get a tour by id
// @Tags game,tour
// @Accept plain
// @Produce json
// @Param id path int true "id of the tour"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/tours/{id} [get]
func GetTour(c *gin.Context) {
	appG := app.Gin{C: c}

	switch tour, err := gameService.GetTour(c.Param("id")); err {

	case C.ErrTourIDNotNumber:
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error())

	case C.ErrTourNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_TOUR_NO_RECORD, err.Error())

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_TOUR_FAIL, err.Error())

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, tour)

	default:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err)
	}
}

// GetTotalTours godoc
// @Summary Get total num of tours
// @Description Get total num of tours
// @Tags game,tour
// @Accept plain
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/tours [get]
func GetTotalTours(c *gin.Context) {
	appG := app.Gin{C: c}

	switch total, err := gameService.GetTotalTours(); err {

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error())

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, total)

	}
}

// AddTour godoc
// @Summary Add a new tour
// @Description Add a new tour
// @Tags game,tour
// @Accept json
// @Produce json
// @Param token header string true "auth token, must register & login to get the token"
// @Param songs body string true "id of the song, should have many"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 404 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /game/tours/new [post]
func AddTour(c *gin.Context) {
	appG := app.Gin{C: c}

	var t tour
	if err := c.BindJSON(&t); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error())
		return
	}

	switch tourID, err := gameService.AddTour(t.Collects); err {

	case C.ErrTourAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_TOUR_FORMAT_INCORRECT, err.Error())

	case C.ErrTourAddCollectsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_TOUR_NO_COLLECTS_RECORD, err.Error())

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error())

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, tourID)

	}
}
