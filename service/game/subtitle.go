package game

import (
	"database/sql"
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/constant"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
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

func (srv *Service) ListYoutubeCaptionLanguages(param string) (map[string]string, error) {
	youtube := subtitle.NewWebSubtitleFactory("youtube")

	languages, err := youtube.ListLanguages(param)
	if err != nil {
		return nil, constant.ErrCaptionError // TODO: error described more detailed
	}
	return languages, nil
}

func (srv *Service) ConvertFileToSubtitle(filetype string, file []byte) ([]model.Lyric, error) {
	if filetype != "srt" && filetype != "lrc" {
		return nil, C.ErrConvertFileToSubtitleTypeNotSupported
	}
	if len(file) == 0 {
		return nil, C.ErrConvertFileToSubtiteParse
	}
	lines, err := subtitle.NewSubtitleFactory(filetype).ReadFromBytes(file)
	if err != nil {
		log.WarnWithSource("ConvertFileToSubtitle: parse lyrics error: ", err)
		return nil, C.ErrConvertFileToSubtiteParse
	}
	return model.ParseLyrics(lines), nil
}

func (srv *Service) DownloadYoutubeSubtitle(url string, languageCode string) ([]model.Lyric, error) {

	lines, err := subtitle.NewWebSubtitleFactory("youtube").GetLines(url, languageCode)
	if err != nil {
		return nil, C.ErrDownloadYoutubeSubtitle // TODO: error described more detailed
	}

	lyrics := model.ParseLyrics(lines)
	return lyrics, nil
}
