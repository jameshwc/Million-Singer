package subtitle

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/log"
)

func parseLrcDuration(s string) (time.Duration, error) {
	return parseDuration("00:"+s, ".", 2)
}

func ReadLrcFromFile(i io.Reader) ([]model.Lyric, error) {
	scanner := bufio.NewScanner(i)
	var lyrics []model.Lyric
	var prev *model.Lyric = nil
	idx := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if match, err := regexp.Match("^\\[\\d+:\\d+\\.\\d+\\]", []byte(line)); err != nil {
			log.Error(err)
			return nil, err
		} else if !match {
			continue
		}
		var lyric model.Lyric
		splits := strings.Split(line, "]")
		duration, err := parseLrcDuration(splits[0][1:])
		if err != nil {
			log.Error(err)
			return nil, err
		}
		lyric.Line = splits[1]
		lyric.StartAt = duration
		lyric.Index = idx
		if prev != nil {
			prev.EndAt = duration
			lyrics = append(lyrics, *prev)
		}
		prev = &lyric
		idx++
	}
	prev.EndAt = time.Duration(int64(1) << 62) // infinite end_at for the last lyric
	lyrics = append(lyrics, *prev)
	fmt.Println(lyrics)
	return lyrics, nil
}

func ReadLrcFromBytes(file []byte) ([]model.Lyric, error) {
	return ReadLrcFromFile(bytes.NewReader(file))
}
