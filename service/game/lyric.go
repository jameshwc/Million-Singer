package game

import (
	"database/sql"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/repo"
)

func GetLyricsWithSongID(param string) ([]*model.Lyric, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, constant.ErrSongIDNotNumber
	}
	s, err := repo.Song.Get(id, true)
	if err == sql.ErrNoRows {
		return nil, constant.ErrSongNotFound
	}
	return s.Lyrics, nil
}
