package model

import (
	"math/rand"
	"strconv"
	"strings"
)

type Song struct {
	ID         int      `json:"id,omitempty"`
	Lyrics     []*Lyric `json:"lyrics,omitempty"`
	VideoID    string   `json:"video_id,omitempty"`
	StartTime  string   `json:"start_time,omitempty"`
	EndTime    string   `json:"end_time,omitempty"`
	Language   string   `json:"language,omitempty"`
	Name       string   `json:"name"`
	Singer     string   `json:"singer"`
	Genre      string   `json:"genre,omitempty"`
	MissLyrics string   `json:"miss_lyrics,omitempty"` // IDs (integers) with comma seperated
}

func (s *Song) RandomGetMissLyricID() int {
	numsStrs := strings.Split(s.MissLyrics, ",")
	// Let's assume MissLyrics had been validated so we can skip error check
	randID, _ := strconv.Atoi(numsStrs[rand.Intn(len(numsStrs))])
	return randID
}
