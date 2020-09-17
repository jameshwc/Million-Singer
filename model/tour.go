package model

import (
	"log"

	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model `json:"-"`
	Levels     []*Level `gorm:"many2many:tour_songs;" json:"levels"`
}

func GetTour(id int) (*Tour, error) {
	var tour Tour
	stmt := db.Session(&gorm.Session{DryRun: true})
	log.Println(stmt.Preload("Levels").Where("id = ?", id).First(&tour).Statement.SQL.String())
	err := db.Preload("Levels").Where("id = ?", id).First(&tour).Error
	if err != nil {
		return nil, err
	}
	return &tour, nil
}

func GetTotalTours() (int64, error) {
	var tours []*Tour
	rows := db.Find(&tours)
	if rows.Error != nil {
		return 0, rows.Error
	}
	return rows.RowsAffected, nil
}

func (t *Tour) Commit() error {
	return db.Create(t).Error
}
