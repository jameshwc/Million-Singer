package repo

import (
	"fmt"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo/mysql"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jameshwc/Million-Singer/conf"
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

func Setup() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBconfig.User,
		conf.DBconfig.Password,
		conf.DBconfig.Host,
		conf.DBconfig.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	Song = mysql.NewSongRepository(db)
	Collect = mysql.NewCollectRepository(db)
	Tour = mysql.NewTourRepository(db)
	User = mysql.NewUserRepository(db)
	Cache = gredis.NewRedisRepository()
}
