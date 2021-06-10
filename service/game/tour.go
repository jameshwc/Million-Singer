package game

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"
	"github.com/jameshwc/Million-Singer/service/cache"
)

func (srv *Service) GetTour(param string) (*model.Tour, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Debugf("Get Tour: param id %s is not a number", param)
		return nil, C.ErrTourIDNotNumber
	}

	key := cache.GetTourKey(id)
	if data, err := repo.Cache.Get(key); err == nil {
		var t model.Tour
		if err := json.Unmarshal(data, &t); err == nil {
			log.Info("Get Tour: redis being used to get tour")
			return &t, nil
		}
		log.Info("Get Tour: unable to unmarshal data: ", err)
	}

	tour, err := repo.Tour.Get(id)
	if err == sql.ErrNoRows {
		log.Debugf("Get Tour: tour id %d record not found", id)
		return nil, C.ErrTourNotFound
	} else if err != nil {
		log.Error("Get Tour: unknown database error", err.Error())
		return nil, C.ErrDatabase
	}
	repo.Cache.Set(key, tour, 7200)
	return tour, nil
}

func (srv *Service) GetTotalTours() (int, error) {
	total, err := repo.Tour.GetTotal()
	if err != nil {
		log.Error("Get Total Tours: unknown database error, ", err.Error())
		return 0, C.ErrDatabase
	}
	return total, nil
}

func (srv *Service) AddTour(collectsID []int, title string) (int, error) {

	if len(collectsID) == 0 || title == "" {
		return 0, C.ErrTourAddFormatIncorrect
	}

	if checkDuplicateInts(collectsID) {
		return 0, C.ErrTourAddCollectsDuplicate
	}

	collectNum, err := repo.Collect.CheckManyExist(collectsID)
	if err != nil {
		log.Error("Check Collects Exist: ", err)
		return 0, C.ErrDatabase
	} else if len(collectsID) != int(collectNum) {
		return 0, C.ErrTourAddCollectsRecordNotFound
	}
	id, err := repo.Tour.Add(collectsID, title)
	if err != nil {
		log.Error("Add Tour: ", err)
		return 0, C.ErrDatabase
	}

	return id, nil
}

func (srv *Service) DelTour(param string) error {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Debugf("Del Tour: param id %s is not a number", param)
		return C.ErrTourIDNotNumber
	}
	total, err := repo.Tour.GetTotal()
	if err != nil {
		return C.ErrDatabase
	}

	if id < 0 || id >= total {
		return C.ErrTourDelIDIncorrect
	}

	if tour, _ := repo.Tour.Get(id); tour == nil {
		return C.ErrTourDelDeleted
	}

	if err := repo.Tour.Del(id); err != nil {
		log.Error("Del Tour: ", err)
		return C.ErrDatabase
	}

	key := cache.GetTourKey(id)
	if err = repo.Cache.Del(key); err != nil {
		log.WarnWithSource(err)
	}

	return nil
}
