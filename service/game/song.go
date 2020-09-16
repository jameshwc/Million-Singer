package game

import (
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
	"gorm.io/gorm"
)

func lyricsJoin(lyrics []int) string {
	s := make([]string, len(lyrics))
	for i, v := range lyrics {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func findMax(l []int) (int, error) {
	var max int
	for i := range l {
		if l[i] > max {
			max = l[i]
		}
		if l[i] < 0 {
			return 0, errors.New("lyrics id negative")
		}
	}
	return max, nil
}

type Song struct {
	File       []byte `json:"file" valid:"Required;"`
	FileType   string `json:"file_type" valid:"Required;MaxSize(50)"`
	URL        string `json:"url" valid:"Required;Match(/^https?://)"`
	Name       string `json:"name" valid:"Required;MaxSize(100)"`
	Singer     string `json:"singer" valid:"Required;MaxSize(100)"`
	MissLyrics []int  `json:"miss_lyrics" valid:"Required;"`
	Genre      string `json:"genre"`
	Language   string `json:"language"`
}

type SongInstance struct {
	Song        *model.Song
	MissLyricID int `json:"miss_lyric_id"`
}

func AddSong(s *Song) (uint, error) {

	valid := validation.Validation{}

	ok, _ := valid.Valid(s)
	if !ok {
		return 0, C.ErrSongFormatIncorrect
	}

	var song model.Song
	var err error
	switch s.FileType {
	case "srt":
		song.Lyrics, err = subtitle.ReadSrtFromBytes(s.File)
	default:
		return 0, C.ErrSongLyricsFileTypeNotSupported
	}
	if err != nil {
		return 0, C.ErrSongParseLyrics
	}

	maxIdx, err := findMax(s.MissLyrics)
	if err != nil || maxIdx > len(song.Lyrics) {
		return 0, C.ErrSongMissLyricsIncorrect
	}

	song.MissLyrics = lyricsJoin(s.MissLyrics)
	song.Genre = s.Genre
	song.Language = s.Language
	song.URL = s.URL
	song.Singer = s.Singer
	song.Name = s.Name

	if err = song.Commit(); err != nil {
		return 0, C.ErrDatabase
	}

	return song.ID, nil
}

func GetSongInstance(param string, hasLyrics bool) (*SongInstance, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrSongIDNotNumber
	}
	s, err := model.GetSong(id, hasLyrics)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrSongNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	return &SongInstance{Song: s, MissLyricID: s.RandomGetMissLyricID()}, nil
}
