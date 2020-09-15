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

func GetTotalTours() (int64, error) {
	total, err := model.GetTotalTours()
	if err != nil {
		return 0, C.ErrDatabase
	}
	return total, nil
}
