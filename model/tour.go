package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model `json:"-"`
	LevelID    []Level `gorm:"many2many:game_levels;" json:"levels"`
}

func GetTour(id int) (*Tour, error) {
	var tour Tour
	err := db.Where("id = ?", id).First(&tour).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tour, nil
}