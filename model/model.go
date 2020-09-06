package model

import (
	"fmt"
	"log"

	"github.com/jameshwc/Million-Singer/conf"
	model "github.com/jameshwc/Million-Singer/model/game"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Conf.DB.User,
		conf.Conf.DB.Password,
		conf.Conf.DB.Host,
		conf.Conf.DB.Name)), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.AutoMigrate(&model.Game{}, &model.Level{}, &model.GameSong{}, &model.Song{}, &model.Lyric{})
	db.Create(&model.Game{LevelsID: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, GameID: 0})
}
