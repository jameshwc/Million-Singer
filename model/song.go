package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	// SongID     int     `json:"song_id"`
	Lyrics     []Lyric `json:"lyrics"`
	URL        string  `json:"url"`
	StartTime  string  `json:"start_time"`
	EndTime    string  `json:"end_time"`
	Language   string  `json:"language"`
	Name       string  `json:"name"`
	Singer     string  `json:"singer"`
	Genre      string  `json:"genre"`
	MissLyrics string  // IDs (integers) with comma seperated
}

type SongInstance struct {
	S           Song
	MissLyricID int `json:"miss_lyric_id"`
}

func AddSong(attr map[string]string, lyrics []Lyric) {

}

func GetSongInstance(SongID int) (*SongInstance, error) {
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
