package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/pkg/app"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	gameService "github.com/jameshwc/Million-Singer/service/game"
)

// GetTour godoc
// @Summary Add a new tour
// @Description Add a new tour
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
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, nil)

	case C.ErrTourNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_TOUR_NO_RECORD, nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_TOUR_FAIL, nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, tour)

	}
}
