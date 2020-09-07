package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	LevelID []Level `gorm:"many2many:game_levels;" json:"levels"`
	GameID  int     `gorm:"primaryKey" json:"game_id"`
}

func GetGame(id int) (*Game, error) {
	var game Game
	err := db.Where("id = ?", id).First(&game).Error
	fmt.Println(err)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &game, nil
}
