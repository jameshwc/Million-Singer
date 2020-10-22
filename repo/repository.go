package repo

import (
	"github.com/jameshwc/Million-Singer/model"

	_ "github.com/go-sql-driver/mysql"
)

type TourRepo interface {
	Add(collectsID []int) (int, error)
	Get(id int) (*model.Tour, error)
	GetTotal() (int, error)
}

type CollectRepo interface {
	Add(title string, songsID []int) (int, error)
	Get(id int) (*model.Collect, error)
	CheckManyExist(collectsID []int) (int64, error)
}
type SongRepo interface {
	Add(videoID, name, singer, genre, language, missLyrics, startTime, endTime string, lyrics []model.Lyric) (int, error)
	Get(songID int, hasLyrics bool) (*model.Song, error)
	Delete(id int) error
	QueryByVideoID(videoID string) (id int64, err error)
	CheckManyExist(songsID []int) (int64, error)
}
type UserRepo interface {
	Auth(username string, password string) (*model.User, error)
	Add(name, email, password string) (int64, error)
	IsEmailDuplicate(email string) bool
	IsNameDuplicate(name string) bool
}

type CacheRepo interface {
	Set(key string, data interface{}, timeout int) error
	Get(key string) ([]byte, error)
	Del(key string) error
}

var Tour TourRepo
var Collect CollectRepo
var Song SongRepo
var User UserRepo
var Cache CacheRepo
