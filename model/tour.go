package model

import (
	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model `json:"-"`
	Collects   []*Collect `gorm:"many2many:tour_collects;" json:"collects"`
}

func GetTour(id int) (*Tour, error) {
	var tour Tour
	err := db.Preload("Collects").Where("id = ?", id).First(&tour).Error
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
