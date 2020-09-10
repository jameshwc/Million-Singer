package model

import (
	"time"

	"gorm.io/gorm"
)

type Lyric struct {
	gorm.Model `json:"-"`
	Index      int           `json:"index"`
	Line       string        `sql:"type:VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8_general_ci" json:"line"`
	StartAt    time.Duration `json:"start_time"`
	EndAt      time.Duration `json:"end_time"`
	SongID     uint          `json:"-"`
}
