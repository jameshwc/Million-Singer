package game

import (
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"

	"github.com/jameshwc/Million-Singer/pkg/subtitle"
	"github.com/jameshwc/Million-Singer/service/cache"
)

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
	*model.Song
	MissLyricID int `json:"miss_lyric_id"`
}

func (srv *Service) AddSong(s *Song) (int, error) {

	valid := validation.Validation{}

	ok, err := valid.Valid(s)
	if !ok || (s.FileType != "youtube" && len(s.File) == 0) {
		if !ok && err != nil {
			log.Error("Add Song: valid.Valid not ok, error: ", err.Error())
		} else if !ok {
			log.Info("Add Song: valid.Valid not ok")
		} else {
			log.Info("Add Song: file column not offer and file type not youtube")
		}
		return 0, C.ErrSongFormatIncorrect
	}

	if checkDuplicateInts(s.MissLyrics) {
		return 0, C.ErrSongAddLyricsIndexDuplicate
	}

	videoID, err := subtitle.ParseVideoID(s.URL)
	if err != nil {
		log.Debug("Add Song: parse video id error: ", err.Error())
		return 0, C.ErrSongAddURLIncorrect
	}

	switch id, err := repo.Song.QueryByVideoID(videoID); err {
	case sql.ErrNoRows:
		break
	case nil:
		log.Debug("Add Song: add duplicate song", int(id))
		return int(id), C.ErrSongAddDuplicate
	default:
		log.Error("Add Song: unknown database error: ", err.Error())
		return 0, C.ErrDatabase
	}

	var lines []subtitle.Line
	switch s.FileType {
	case "srt", "lrc":
		lines, err = subtitle.NewSubtitleFactory(s.FileType).ReadFromBytes(s.File)
	case "youtube":
		lines, err = subtitle.NewWebSubtitleFactory(s.FileType).GetLines(s.URL, s.Language)
	default:
		log.Debug("Add Song: file type not supported: ", s.FileType)
		return 0, C.ErrSongAddLyricsFileTypeNotSupported
	}

	if err != nil {
		log.WarnWithSource("Add Song: parse lyrics error: ", err)
		return 0, C.ErrSongAddParseLyrics
	}

	lyrics := model.ParseLyrics(lines)

	maxIdx := findMax(s.MissLyrics)
	if maxIdx < 0 || maxIdx > len(lyrics) {
		log.Info("Add Song: miss lyrics id out of index or negative")
		return 0, C.ErrSongAddMissLyricsIncorrect
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

func (srv *Service) GetSongs() ([]*model.Song, error) {
	// TODO: Cache?
	songs, err := repo.Song.Gets()
	if err != nil {
		return nil, C.ErrDatabase
	}
	return songs, nil
}
