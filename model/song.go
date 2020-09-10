package model

import (
	"math/rand"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model `json:"-"`
	FrontendID uint     `json:"id"`
	Lyrics     []Lyric  `json:"lyrics"`
	URL        string   `json:"url"`
	StartTime  string   `json:"start_time"`
	EndTime    string   `json:"end_time"`
	Language   string   `json:"language"`
	Name       string   `json:"name"`
	Singer     string   `json:"singer"`
	Genre      string   `json:"genre"`
	MissLyrics string   `json:"-"` // IDs (integers) with comma seperated
	Levels     []*Level `gorm:"many2many:level_songs;" json:"-"`
}

func (s *Song) Commit() error {
	if err := db.Create(s).Error; err != nil {
		return err
	}
	if err := db.Model(s).UpdateColumn("FrontendID", s.ID).Error; err != nil {
		return err
	}
	return nil
}

func (s *Song) RandomGetMissLyricID() int {
	numsStrs := strings.Split(s.MissLyrics, ",")
	// Let's assume MissLyrics had been validated so we can skip error check
	randID, _ := strconv.Atoi(numsStrs[rand.Intn(len(numsStrs))])
	return randID
}

func GetSong(songID int, hasLyrics bool) (*Song, error) {
	var song Song
	var subdb *gorm.DB
	if hasLyrics {
		subdb = db.Preload("Lyrics")
	} else {
		subdb = db
	}
	if err := subdb.Where("id = ?", songID).First(&song).Error; err != nil {
		return nil, err
	}
	// db.Where("song_id = ?", songID).Find(&song.Lyrics)
	song.FrontendID = song.ID // workaround; TODO: find a pretty solution
	return &song, nil
}

func GetSongs(songsID []int) ([]*Song, error) {
	var songs []*Song
	err := db.Find(&songs, songsID).Error
	if err != nil {
		return nil, err
	}
	for i := range songs {
		songs[i].FrontendID = songs[i].ID
	}
	return songs, nil
}
