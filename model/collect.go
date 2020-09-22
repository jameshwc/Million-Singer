package model

import (
	"errors"

	"gorm.io/gorm"
)

type Collect struct {
	gorm.Model `json:"-"`
	Title      string  `json:"title"`
	Songs      []*Song `gorm:"many2many:collect_songs;" json:"songs"`
	FrontendID uint    `json:"id"`
}

func (l *Collect) Commit() error {
	if err := db.Create(l).Error; err != nil {
		return err
	}
	if err := db.Model(l).UpdateColumn("FrontendID", l.ID).Error; err != nil {
		return err
	}
	return nil
}

func GetCollect(collectID int) (*Collect, error) {
	var collect Collect
	if err := db.Preload("Songs").Where("id = ?", collectID).First(&collect).Error; err != nil {
		return nil, err
	}
	collect.FrontendID = collect.ID
	return &collect, nil
}

func GetCollects(collectsID []int) ([]*Collect, error) {
	var collects []*Collect
	err := db.Find(&collects, collectsID).Error
	if err != nil {
		return nil, err
	}
	if len(collects) != len(collectsID) {
		return nil, errors.New("some collects ID are incorrect")
	}
	for i := range collects {
		collects[i].FrontendID = collects[i].ID
	}
	return collects, err
}
