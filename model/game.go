package model

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	LevelsID []int `json: "levels_id"`
	GameID   int   `json: "game_id`
}

func GetGame(id int) (*Game, error) {
	var game Game
	err := db.Where("game_id = ?", id).First(&game).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &game, nil
}
