package model

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	LevelID  []Level `gorm:"many2many:game_levels;" json:"levels"`
	LevelsID string
	GameID   int `gorm:"primaryKey" json:"game_id"`
}

func GetGame(id int) (*Game, error) {
	var game Game
	err := db.Where("game_id = ?", id).First(&game).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &game, nil
}
