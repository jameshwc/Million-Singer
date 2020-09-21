package model

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jameshwc/Million-Singer/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Setup(externalDB *gorm.DB) {
	var err error
	db = externalDB
	if db == nil {
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBconfig.User,
			conf.DBconfig.Password,
			conf.DBconfig.Host,
			conf.DBconfig.Name)), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatalf("models.Setup err: %v", err)
		}
	}
	// db.AutoMigrate(&Tour{}, &Level{}, &Lyric{}, &Song{}, &User{})
	db = db.Debug()
}
