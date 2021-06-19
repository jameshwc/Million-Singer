package subtitle

import (
	"io"
	"time"

	"github.com/jameshwc/Million-Singer/pkg/log"
)

type Subtitler interface {
	ReadFromFile(i io.Reader) ([]Line, error)
	ReadFromBytes(file []byte) ([]Line, error)
}

type Weber interface {
	GetLines(url, languageCode string) ([]Line, error)
	ListLanguages(url string) (map[string]string, error)
	GetTitle(url string) (string, error)
}

type Line struct {
	Index   int
	Text    string
	StartAt time.Duration
	EndAt   time.Duration
}

var ( // workaround for easier unit tests: dependency injection
	SRT     Subtitler = &srt{}
	LRC     Subtitler = &lrc{}
	Youtube Weber     = &youtube{}
)

func NewSubtitleFactory(filetype string) Subtitler {
	switch filetype {
	case "srt":
		return SRT
	case "lrc":
		return LRC
	default:
		log.Error("not supported subtitle extension")
		return nil
	}
}

func NewWebSubtitleFactory(filetype string) Weber {
	switch filetype {
	case "youtube":
		return Youtube
	default:
		log.Error("not supported web subtitle filetype")
		return nil
	}
}
