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

func GetCollect(param string) (*model.Collect, error) {

	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrCollectIDNotNumber
	}

	key := cache.GetCollectKey(id)
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			log.Info(err)
		} else {
			log.Info("redis being used to get collect")
			var c model.Collect
			if err := json.Unmarshal(data, &c); err != nil {
				log.Info("unable to unmarshal data: ", err)
			} else {
				return &c, nil
			}
		}
	}

	collect, err := model.GetCollect(id)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrCollectNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	return collect, nil
}

func AddCollect(songs []int, title string) (uint, error) {
	if len(songs) == 0 || title == "" {
		return 0, C.ErrCollectAddFormatIncorrect
	}

	var l model.Collect
	var err error

	l.Songs, err = model.GetSongs(songs)
	if err != nil {
		return 0, C.ErrCollectAddSongsRecordNotFound
	}

	l.Title = title
	if err = l.Commit(); err != nil {
		return 0, C.ErrDatabase
	}

	return l.ID, nil
}
