package game

import (
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"gorm.io/gorm"
)

func GetLevel(param string) (*model.Level, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrLevelIDNotNumber
	}
	level, err := model.GetLevel(id)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrLevelNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	return level, nil
}

func AddLevel(songs []int, title string) (uint, error) {
	if len(songs) == 0 || title == "" {
		return 0, C.ErrLevelAddFormatIncorrect
	}

	var l model.Level
	var err error

	l.Songs, err = model.GetSongs(songs)
	if err != nil {
		return 0, C.ErrLevelAddSongsRecordNotFound
	}

	l.Title = title
	if err = l.Commit(); err != nil {
		return 0, C.ErrDatabase
	}

	return l.ID, nil
}
