package model

import (
	"time"

	"github.com/jameshwc/Million-Singer/pkg/subtitle"
)

type Lyric struct {
	Index   int           `json:"index"`
	Line    string        `sql:"type:VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8_general_ci" json:"line"`
	StartAt time.Duration `json:"start_time"`
	EndAt   time.Duration `json:"end_time"`
}

func ParseLyrics(lines []subtitle.Line) (lyrics []Lyric) {
	for _, line := range lines {
		lyrics = append(lyrics, Lyric{line.Index, line.Text, time.Duration(line.StartAt.Milliseconds()), time.Duration(line.EndAt.Milliseconds())})
	}
	return
}
