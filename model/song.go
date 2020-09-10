package model

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Song struct {
	gorm.Model `json:"-"`
	FrontendID uint    `json:"id"`
	Lyrics     []Lyric `json:"lyrics"`
	URL        string  `json:"url"`
	StartTime  string  `json:"start_time"`
	EndTime    string  `json:"end_time"`
	Language   string  `json:"language"`
	Name       string  `json:"name"`
	Singer     string  `json:"singer"`
	Genre      string  `json:"genre"`
	MissLyrics string  `json:"-"` // IDs (integers) with comma seperated
}

func (s *Song) Commit() error {
	if err := db.Create(s).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *Song) RandomGetMissLyricID() int {
	numsStrs := strings.Split(s.MissLyrics, ",")
	nums := make([]int, len(numsStrs))
	// Let's assume MissLyrics had been validated so we can skip error check
	for i := range numsStrs {
		nums[i], _ = strconv.Atoi(numsStrs[i])
	}
	return nums[rand.Intn(len(nums))]
}
func GetSong(SongID int) (*Song, error) {
	var song Song
	err := db.Where("id = ?", SongID).First(&song).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	db.Where("song_id = ?", SongID).Find(&song.Lyrics)
	song.FrontendID = song.ID // workaround; TODO: find a pretty solution
	return &song, nil
}
