package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title   string     `json:"title"`
	Songs   []GameSong `gorm:"many2many:level_gamesongs" json:"songs"`
	LevelID int        `json:"level_id"`
}
