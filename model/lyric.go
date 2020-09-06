package model

import "gorm.io/gorm"

type Lyric struct {
	gorm.Model
	Line    string `json:"line"`
	LyricID int    `json:"lyric_id"`
	Start   string `json:"start_time"`
	End     string `json:"end_time"`
	SongID  string `json:"song_id"`
}

func GetLyric(id int) {

}
