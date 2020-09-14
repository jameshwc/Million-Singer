package game

import (
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"gorm.io/gorm"
)

func GetTour(param string) (*model.Tour, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrTourIDNotNumber
	}
	tour, err := model.GetTour(id)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrTourNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	return tour, nil
}
