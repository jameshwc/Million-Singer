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

func GetLyricsWithSongID(songID int) (lyrics []*Lyric, err error) {
	if err = db.Where("song_id = ?", songID).Find(&lyrics).Error; err != nil {
		return nil, err
	}
	return
}
