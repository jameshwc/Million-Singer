package game

import (
	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
)

func (srv *Service) ConvertFileToSubtitle(filetype string, file []byte) ([]model.Lyric, error) {
	if filetype != "src" && filetype != "lrc" {
		return nil, C.ErrConvertFileToSubtitleTypeNotSupported
	}
	lines, err := subtitle.NewSubtitleFactory(filetype).ReadFromBytes(file)
	if err != nil {
		log.WarnWithSource("ConvertFileToSubtitle: parse lyrics error: ", err)
		return nil, C.ErrConvertFileToSubtiteParse
	}
	return model.ParseLyrics(lines), nil
}
