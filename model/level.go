package model

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model `json:"-"`
	Title      string  `json:"title"`
	Songs      []*Song `gorm:"many2many:level_song;" json:"songs"`
	// SongsID    []int   `gorm:"-" json:"songs_id"`
}

func (l *Level) Commit() error {
	return db.Create(l).Error
}

func GetLevel(levelID int) (*Level, error) {
	var level Level
	if err := db.Preload("Songs").Where("id = ?", levelID).First(&level).Error; err != nil {
		return nil, err
	}
	return &level, nil
}

func GetLevels(levelsID []int) ([]*Level, error) {
	var levels []*Level
	err := db.Find(&levels, levelsID).Error
	return levels, err
}
