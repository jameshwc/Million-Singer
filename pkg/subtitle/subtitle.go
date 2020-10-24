package subtitle

import (
	"io"
	"time"
)

type Subtitler interface {
	Srter
	Lrcer
	Youtuber
}

type Srter interface {
	ReadSrtFromFile(i io.Reader) ([]Line, error)
	ReadSrtFromBytes(file []byte) ([]Line, error)
}

type Lrcer interface {
	ReadLrcFromFile(i io.Reader) ([]Line, error)
	ReadLrcFromBytes(file []byte) ([]Line, error)
}

type Youtuber interface {
	GetLines(url string) ([]Line, error)
}

type Line struct {
	Index   int
	Text    string
	StartAt time.Duration
	EndAt   time.Duration
}

type Subtitle struct{}
type Srt struct{}
type Lrc struct{}
type Youtube struct{}

var subtitle = &Subtitle{}
var srt = &Srt{}
var lrc = &Lrc{}
var youtube = &Youtube{}

func (y *Youtube) GetLines(url string) ([]Line, error) {
	youtube, err := newYoutubeDownloader(url)
	if err != nil {
		return nil, err
	}

	return youtube.getLyrics()
}

func GetLyricsFromYoutubeSubtitle(url string) ([]Line, error) {
	return youtube.GetLines(url)
}
