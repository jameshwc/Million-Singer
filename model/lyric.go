package model

import (
	"time"

	"gorm.io/gorm"
)

type Lyric struct {
	gorm.Model
	Index   int           `json:"index"`
	Line    string        `gorm:"type:utf8mb4" json:"line"`
	StartAt time.Duration `json:"start_time"`
	EndAt   time.Duration `json:"end_time"`
	SongID  uint          `json:"song_id"`
}
