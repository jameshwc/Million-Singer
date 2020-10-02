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

func GetCollect(param string) (*model.Collect, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrCollectIDNotNumber
	}

	key := cache.GetCollectKey(id)
	if data, err := gredis.Get(key); err == nil {
		log.Info("redis being used to get collect")
		var c model.Collect
		if err := json.Unmarshal(data, &c); err != nil {
			log.Info("unable to unmarshal data: ", err)
		} else {
			return &c, nil
		}
	}

	collect, err := model.GetCollect(id)
	if err == sql.ErrNoRows {
		return nil, C.ErrCollectNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	gredis.Set(key, collect, 7200)
	return collect, nil
}

func AddCollect(songs []int, title string) (int, error) {
	if len(songs) == 0 || title == "" {
		return 0, C.ErrCollectAddFormatIncorrect
	}

	songNum, err := model.CheckSongsExist(songs)
	if err != nil {
		log.Error("Check songs Exist: ", err)
		return 0, C.ErrDatabase
	} else if len(songs) != int(songNum) {
		return 0, C.ErrCollectAddSongsRecordNotFound
	}
	id, err := model.AddCollect(title, songs)
	if err != nil {
		return 0, C.ErrDatabase
	}
	return id, nil

}
