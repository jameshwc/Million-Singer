package model

import (
	"fmt"
	"log"

	"github.com/jameshwc/Million-Singer/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBconfig.User,
		conf.DBconfig.Password,
		conf.DBconfig.Host,
		conf.DBconfig.Name)), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.AutoMigrate(&Game{}, &Lyric{}, &Song{})
}
