package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title   string `json:"title"`
	Songs   []Song `gorm:"many2many:level_songs" json:"songs"`
	LevelID int    `json:"level_id"`
}
