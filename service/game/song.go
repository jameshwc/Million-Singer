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
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"

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

func (srv *Service) AddSong(s *Song) (int, error) {

	valid := validation.Validation{}

	ok, err := valid.Valid(s)
	if !ok || (s.FileType != "youtube" && len(s.File) == 0) {
		if !ok {
			log.Debug("Add Song: valid.Valid not ok, ", err.Error())
		} else {
			log.Debug("Add Song: file column not offer and file type not youtube")
		}
		return 0, C.ErrSongFormatIncorrect
	}

	videoID, err := subtitle.ParseVideoID(s.URL)
	if err != nil {
		log.Debug("Add Song: parse video id error: ", err.Error())
		return 0, C.ErrSongURLIncorrect
	}

	switch id, err := repo.Song.QueryByVideoID(videoID); err {
	case sql.ErrNoRows:
		break
	case nil:
		log.Debug("Add Song: add duplicate song", int(id))
		return int(id), C.ErrSongDuplicate
	default:
		log.Error("Add Song: unknown database error: ", err.Error())
		return 0, C.ErrDatabase
	}

	var lyrics []model.Lyric
	switch s.FileType {
	case "srt":
		lyrics, err = subtitle.ReadSrtFromBytes(s.File)
	case "lrc":
		lyrics, err = subtitle.ReadLrcFromBytes(s.File)
	case "youtube":
		lyrics, err = subtitle.GetLyricsFromYoutubeSubtitle(s.URL)
	default:
		log.Debug("Add Song: file type not supported: ", s.FileType)
		return 0, C.ErrSongLyricsFileTypeNotSupported
	}

	if err != nil {
		log.WarnWithSource("Add Song: parse lyrics error: ", err)
		return 0, C.ErrSongParseLyrics
	}

	maxIdx, err := findMax(s.MissLyrics)
	if err != nil || maxIdx > len(lyrics) {
		if err != nil {
			log.Info("Add Song: find miss lyrics index error: ", err.Error())
		} else {
			log.Info("Add Song: miss lyrics id out of index")
		}
		return 0, C.ErrSongMissLyricsIncorrect
	}

	id, err := repo.Song.Add(videoID, s.Name, s.Singer, s.Genre, s.Language, lyricsJoin(s.MissLyrics), "", "", lyrics)
	if err != nil {
		log.Error("Add Song: unknown database error: ", err.Error())
		return 0, C.ErrDatabase
	}

	return id, nil
}

func (srv *Service) GetSongInstance(param string, hasLyrics bool) (*SongInstance, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Debugf("Get Song: param id %s is not a number", param)
		return nil, C.ErrSongIDNotNumber
	}

	key := cache.GetSongKey(id, hasLyrics)
	if data, err := repo.Cache.Get(key); err == nil {
		var s model.Song
		if err := json.Unmarshal(data, &s); err != nil {
			log.Info("Get Song: redis unable to unmarshal data: ", err)
		} else {
			log.Info("Get Song: redis being used to get song")
			return &SongInstance{Song: &s, MissLyricID: s.RandomGetMissLyricID()}, nil
		}
	}
	s, err := repo.Song.Get(id, hasLyrics)
	if err == sql.ErrNoRows {
		log.Infof("Get Song: song id %d not found", id)
		return nil, C.ErrSongNotFound
	} else if err != nil {
		log.Error("Get Song: database error ", err.Error())
		return nil, C.ErrDatabase
	}
	repo.Cache.Set(key, s, 7200)
	return &SongInstance{Song: s, MissLyricID: s.RandomGetMissLyricID()}, nil
}

func (srv *Service) DeleteSong(param string) error {
	id, err := strconv.Atoi(param)
	if err != nil {
		return C.ErrSongIDNotNumber
	}
	if err = repo.Song.Delete(id); err != nil {
		log.Error(err)
		return C.ErrDatabase
	}
	key := cache.GetSongKey(id, true)
	if err = repo.Cache.Del(key); err != nil {
		log.WarnWithSource(err)
	}
	key = cache.GetSongKey(id, false)
	if err = repo.Cache.Del(key); err != nil {
		log.WarnWithSource(err)
	}
	return nil
}
