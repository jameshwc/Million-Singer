package subtitle

import (
	"io"
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
	Index int
}
