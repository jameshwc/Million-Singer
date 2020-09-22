package game

import (
	"encoding/json"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/service/cache"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
)

func GetTour(param string) (*model.Tour, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrTourIDNotNumber
	}

	key := cache.GetTourKey(id)
	if data, err := gredis.Get(key); err == nil {
		log.Info("redis being used to get tour")
		var t model.Tour
		if err := json.Unmarshal(data, &t); err != nil {
			log.Info("unable to unmarshal data: ", err)
		} else {
			return &t, nil
		}
	}

	tour, err := model.GetTour(id)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrTourNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	gredis.Set(key, tour, 7200)
	return tour, nil
}

func GetTotalTours() (int64, error) {
	total, err := model.GetTotalTours()
	if err != nil {
		return 0, C.ErrDatabase
	}
	return total, nil
}

func AddTour(collectsID []int) (uint, error) {

	if len(collectsID) == 0 {
		return 0, C.ErrTourAddFormatIncorrect
	}

	collects, err := model.GetCollects(collectsID)
	if err != nil {
		return 0, C.ErrTourAddCollectsRecordNotFound
	}

	var tour model.Tour
	tour.Collects = collects
	if err := tour.Commit(); err != nil {
		return 0, C.ErrDatabase
	}

	return tour.ID, nil
}
