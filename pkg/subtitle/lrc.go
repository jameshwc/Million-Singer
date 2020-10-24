package subtitle

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/jameshwc/Million-Singer/pkg/log"
)

func parseLrcDuration(s string) (time.Duration, error) {
	return parseDuration("00:"+s, ".", 2)
}

func ReadLrcFromFile(i io.Reader) ([]Line, error) {
	scanner := bufio.NewScanner(i)
	var lines []Line
	var prev *Line = nil
	idx := 1
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 {
			continue
		}
		if match, err := regexp.Match("^\\[\\d+:\\d+\\.\\d+\\]", []byte(text)); err != nil {
			log.Error(err)
			return nil, err
		} else if !match {
			continue
		}
		var line Line
		splits := strings.Split(text, "]")
		duration, err := parseLrcDuration(splits[0][1:])
		if err != nil {
			log.Error(err)
			return nil, err
		}
		line.Text = splits[1]
		line.StartAt = duration
		line.Index = idx
		if prev != nil {
			prev.EndAt = duration
			lines = append(lines, *prev)
		}
		prev = &line
		idx++
	}
	prev.EndAt = time.Duration(int64(1) << 62) // infinite end_at for the last line
	lines = append(lines, *prev)
	return lines, nil
}

func ReadLrcFromBytes(file []byte) ([]Line, error) {
	return ReadLrcFromFile(bytes.NewReader(file))
}
