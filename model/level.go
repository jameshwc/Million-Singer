package model

import (
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model `json:"-"`
	Title      string  `json:"title"`
	Songs      []*Song `gorm:"many2many:level_song;" json:"songs"`
	FrontendID uint    `json:"id"`
	// SongsID    []int   `gorm:"-" json:"songs_id"`
}

func (l *Level) Commit() error {
	if err := db.Create(l).Error; err != nil {
		return err
	}
	if err := db.Model(l).UpdateColumn("FrontendID", l.ID).Error; err != nil {
		return err
	}
	return nil
}

func GetLevel(levelID int) (*Level, error) {
	var level Level
	if err := db.Preload("Songs").Where("id = ?", levelID).First(&level).Error; err != nil {
		return nil, err
	}
	level.FrontendID = level.ID
	return &level, nil
}

func GetLevels(levelsID []int) ([]*Level, error) {
	var levels []*Level
	err := db.Find(&levels, levelsID).Error
	for i := range levels {
		levels[i].FrontendID = levels[i].ID
	}
	return levels, err
}
