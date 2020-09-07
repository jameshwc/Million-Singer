package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title   string         `json:"title"`
	Songs   []SongInstance `gorm:"many2many:level_songinstances" json:"songs"`
	LevelID int            `json:"level_id"`
}
