package model

import (
	"time"

	"gorm.io/gorm"
)

type Lyric struct {
	gorm.Model
	Index int `json:"index`

	Line    string        `json:"line"`
	StartAt time.Duration `json:"start_time"`
	EndAt   time.Duration `json:"end_time"`
	SongID  string        `json:"song_id"`
}

func ConvertSrtToLyrics()
