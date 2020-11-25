package game

import (
	"database/sql"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
	"github.com/jameshwc/Million-Singer/repo"
)

func (srv *Service) GetLyricsWithSongID(param string) ([]*model.Lyric, error) {
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

func (srv *Service) ListYoutubeCaptionLanguages(param string) ([]string, error) {
	youtube := subtitle.NewWebSubtitleFactory("youtube")

	languages, err := youtube.ListLanguages(param)
	if err != nil {
		return nil, constant.ErrCaptionError // TODO: error described more detailed
	}
	return languages, nil
}
