package game

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/service/cache"
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
	if err == sql.ErrNoRows {
		return nil, C.ErrTourNotFound
	} else if err != nil {
		log.Error(err)
		return nil, C.ErrDatabase
	}
	gredis.Set(key, tour, 7200)
	return tour, nil
}

func GetTotalTours() (int, error) {
	total, err := model.GetTotalTours()
	if err != nil {
		log.Error(err)
		return 0, C.ErrDatabase
	}
	return total, nil
}

func AddTour(collectsID []int) (int, error) {

	if len(collectsID) == 0 {
		return 0, C.ErrTourAddFormatIncorrect
	}

	collectNum, err := model.CheckCollectsExist(collectsID)
	if err != nil {
		log.Error("Check Collects Exist: ", err)
		return 0, C.ErrDatabase
	} else if len(collectsID) != int(collectNum) {
		return 0, C.ErrTourAddCollectsRecordNotFound
	}
	id, err := model.AddTour(collectsID)
	if err != nil {
		log.Error("Add Tour: ", err)
		return 0, C.ErrDatabase
	}

	return id, nil
}
