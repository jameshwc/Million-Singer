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

func (srv *Service) GetCollect(param string) (*model.Collect, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		log.Infof("Get Collect: param id %s is not a number", param)
		return nil, C.ErrCollectIDNotNumber
	}

	key := cache.GetCollectKey(id)
	if data, err := repo.Cache.Get(key); err == nil {
		var c model.Collect
		if err := json.Unmarshal(data, &c); err == nil {
			log.Debug("Get Collect: redis being used to get collect")
			return &c, nil
		}
		log.Info("Get Collect: unable to unmarshal data: ", err)
	}

	collect, err := repo.Collect.Get(id)
	if err == sql.ErrNoRows {
		log.Debugf("Get Collect: param id %s is not a number", param)
		return nil, C.ErrCollectNotFound
	} else if err != nil {
		log.Error("Get Collect: unknown database error, ", err.Error())
		return nil, C.ErrDatabase
	}
	repo.Cache.Set(key, collect, 7200)
	return collect, nil
}

func (srv *Service) AddCollect(songs []int, title string) (int, error) {
	if len(songs) == 0 || title == "" {
		return 0, C.ErrCollectAddFormatIncorrect
	}

	if checkDuplicateInts(songs) {
		return 0, C.ErrCollectAddSongsDuplicate
	}

	songNum, err := repo.Song.CheckManyExist(songs)
	if err != nil {
		log.Error("Add Collect: ", err.Error())
		return 0, C.ErrDatabase
	} else if len(songs) != int(songNum) {
		log.Debug("Add Collect: Songs record not found, songs id: ", songs)
		return 0, C.ErrCollectAddSongsRecordNotFound
	}
	id, err := repo.Collect.Add(title, songs)
	if err != nil {
		log.Error("Add Collect: unknown database error, ", err.Error())
		return 0, C.ErrDatabase
	}
	return id, nil
}

func (srv *Service) GetCollects() ([]*model.Collect, error) {
	// TODO: Cache?
	cs, err := repo.Collect.Gets()
	if err != nil {
		return nil, C.ErrDatabase
	}
	return cs, nil
}
