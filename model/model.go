package model

import (
	"fmt"
	"log"

	"github.com/jameshwc/Million-Singer/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// type TourDatabase interface {
// 	GetTour(id int) (*Tour, error)
// 	AddTour(*Tour) error
// }

// type LevelDatabase interface {
// 	GetLevel(id int) (*Level, error)
// 	GetLevels(ids []int) ([]*Level, error)
// 	AddLevel(*Level) error
// }

// type SongDatabase interface {
// 	GetSong(id int) (*Song, error)
// 	GetSongs(ids []int) ([]*Song, error)
// }

// type LyricDatabase interface {
// 	FindLyricsWithSongID(id int) ([]*Lyric, error)
// }

// type GameDatabase interface {
// 	TourDatabase
// 	LevelDatabase
// 	SongDatabase
// 	LyricDatabase
// }

func Setup(externalDB *gorm.DB) {
	var err error
	db = externalDB
	if db == nil {
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBconfig.User,
			conf.DBconfig.Password,
			conf.DBconfig.Host,
			conf.DBconfig.Name)), &gorm.Config{})
		if err != nil {
			log.Fatalf("models.Setup err: %v", err)
		}
	}
	db.AutoMigrate(&Tour{}, &Level{}, &Lyric{}, &Song{}, &User{})
}
