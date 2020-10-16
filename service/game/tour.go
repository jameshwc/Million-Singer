package game

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"
	"github.com/jameshwc/Million-Singer/service/cache"
)

func GetTour(param string) (*model.Tour, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Debugf("Get Tour: param id %s is not a number", id)
		return nil, C.ErrTourIDNotNumber
	}

	key := cache.GetTourKey(id)
	if data, err := gredis.Get(key); err == nil {
		var t model.Tour
		if err := json.Unmarshal(data, &t); err != nil {
			log.Info("Get Tour: unable to unmarshal data: ", err)
		} else {
			log.Info("Get Tour: redis being used to get tour")
			return &t, nil
		}
	}

	tour, err := repo.Tour.Get(id)
	if err == sql.ErrNoRows {
		log.Debugf("Get Tour: tour id %d record not found", id)
		return nil, C.ErrTourNotFound
	} else if err != nil {
		log.Error("Get Tour: unknown database error", err.Error())
		return nil, C.ErrDatabase
	}
	gredis.Set(key, tour, 7200)
	return tour, nil
}

func GetTotalTours() (int, error) {
	total, err := repo.Tour.GetTotal()
	if err != nil {
		log.Error("Get Total Tours: unknown database error, ", err.Error())
		return 0, C.ErrDatabase
	}
	return total, nil
}

func AddTour(collectsID []int) (int, error) {

	if len(collectsID) == 0 {
		return 0, C.ErrTourAddFormatIncorrect
	}

	collectNum, err := repo.Collect.CheckManyExist(collectsID)
	if err != nil {
		log.Error("Check Collects Exist: ", err)
		return 0, C.ErrDatabase
	} else if len(collectsID) != int(collectNum) {
		return 0, C.ErrTourAddCollectsRecordNotFound
	}
	id, err := repo.Tour.Add(collectsID)
	if err != nil {
		log.Error("Add Tour: ", err)
		return 0, C.ErrDatabase
	}

	return id, nil
}
