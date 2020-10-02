package model

import (
	"fmt"

	"github.com/jameshwc/Million-Singer/pkg/log"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jameshwc/Million-Singer/conf"
)

var db *sql.DB

func Setup(externalDB *sql.DB) {
	var err error
	db = externalDB
	if db == nil {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.DBconfig.User,
			conf.DBconfig.Password,
			conf.DBconfig.Host,
			conf.DBconfig.Name))
		if err != nil {
			log.Fatalf("models.Setup err: %v", err)
		}
	}
}
