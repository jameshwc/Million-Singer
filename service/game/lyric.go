package game

import (
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"gorm.io/gorm"
)

func GetLyricsWithSongID(param string) (*[]model.Lyric, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, constant.ErrSongIDNotNumber
	}
	s, err := model.GetSong(id, true)
	if err == gorm.ErrRecordNotFound {
		return nil, constant.ErrSongNotFound
	}
	return &s.Lyrics, nil
}
