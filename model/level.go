package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	Title   string `json: "title"`
	SongsID []int  `json: "songs_id"`
	LevelID int    `json: "level_id"`
}
