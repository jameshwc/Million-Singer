package subtitle

import (
	"io"
	"time"
)

type Subtitler interface {
	ReadFromFile(i io.Reader) ([]Line, error)
	ReadFromBytes(file []byte) ([]Line, error)
}

type Weber interface {
	GetLines(url string) ([]Line, error)
}

type Line struct {
	Index   int
	Text    string
	StartAt time.Duration
	EndAt   time.Duration
}

var SRT = newSrt()
var LRC = newLrc()
var Youtube = newYoutube()
