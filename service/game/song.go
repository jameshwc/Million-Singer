package game

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"
	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/pkg/log"

	"github.com/jameshwc/Million-Singer/pkg/subtitle"
	"github.com/jameshwc/Million-Singer/service/cache"
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
	File       []byte `json:"file"`
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

func AddSong(s *Song) (int, error) {

	valid := validation.Validation{}

	ok, _ := valid.Valid(s)
	if !ok || (s.FileType != "youtube" && len(s.File) == 0) {
		return 0, C.ErrSongFormatIncorrect
	}

	switch id, err := model.QuerySongByUrl(s.URL); err {
	case sql.ErrNoRows:
		break
	case nil:
		return int(id), C.ErrSongDuplicate
	default:
		return 0, C.ErrDatabase
	}

	var lyrics []model.Lyric
	var err error
	switch s.FileType {
	case "srt":
		lyrics, err = subtitle.ReadSrtFromBytes(s.File)
	case "lrc":
		lyrics, err = subtitle.ReadLrcFromBytes(s.File)
	case "youtube":
		lyrics, err = subtitle.GetLyricsFromYoutubeSubtitle(s.URL)
	default:
		return 0, C.ErrSongLyricsFileTypeNotSupported
	}

	if err != nil {
		log.WarnWithSource(err)
		return 0, C.ErrSongParseLyrics
	}

	maxIdx, err := findMax(s.MissLyrics)
	if err != nil || maxIdx > len(lyrics) {
		return 0, C.ErrSongMissLyricsIncorrect
	}

	id, err := model.AddSong(s.URL, s.Name, s.Singer, s.Genre, s.Language, lyricsJoin(s.MissLyrics), "", "", lyrics)
	if err != nil {
		log.Error(err)
		return 0, C.ErrDatabase
	}

	return id, nil
}

func GetSongInstance(param string, hasLyrics bool) (*SongInstance, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrSongIDNotNumber
	}

	key := cache.GetSongKey(id, hasLyrics)
	if data, err := gredis.Get(key); err == nil {
		var s model.Song
		log.Info("redis being used to get song")
		if err := json.Unmarshal(data, &s); err != nil {
			log.Info("unable to unmarshal data: ", err)
		} else {
			return &SongInstance{Song: &s, MissLyricID: s.RandomGetMissLyricID()}, nil
		}
	}
	s, err := model.GetSong(id, hasLyrics)
	if err == sql.ErrNoRows {
		return nil, C.ErrSongNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	gredis.Set(key, s, 7200)
	return &SongInstance{Song: s, MissLyricID: s.RandomGetMissLyricID()}, nil
}

// TODO: Delete Song
// func DeleteSong()
