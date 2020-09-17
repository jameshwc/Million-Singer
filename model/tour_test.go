package model

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGetTour(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock fail to new")
		return
	}
	defer db.Close()

	g, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("gorm fail to init")
	}

	Setup(g)

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "level_id"}).AddRow()
}
