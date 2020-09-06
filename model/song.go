package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	SongID    int      `json: "song_id"`
	Lyrics    []Lyric  `json: "lyrics"`
	URL       string   `json: "url"`
	StartTime string   `json: "start_time"`
	EndTime   string   `json: "end_time"`
	Language  string   `json: "language"`
	Name      string   `json: "name"`
	Singer    string   `json: "singer"`
	Genre     []string `json: "genre"`
}

type GameSong struct {
	Song
	GameSongID  int `json: "gamesong_id"`
	MissLyricID int `json: "miss_lyric_id"`
}
