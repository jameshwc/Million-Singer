package subtitle

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	srtTimeSep = " --> "
)

type srt struct{}

func parseSrtDuration(s string) (time.Duration, error) {
	return parseDuration(s, ",", 3)
}

func (s *srt) ReadFromFile(i io.Reader) ([]Line, error) {
	scanner := bufio.NewScanner(i)
	var lines []Line
	var prev string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, srtTimeSep) {
			scanner.Scan()
			var l Line
			var err error
			l.Text = scanner.Text()
			l.Index, err = strconv.Atoi(prev)
			if err != nil {
				return nil, err
			}
			boundaries := strings.Split(line, srtTimeSep)
			if l.StartAt, err = parseSrtDuration(boundaries[0]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[0], err)
			}
			if l.EndAt, err = parseSrtDuration(boundaries[1]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[1], err)
			}
			if l.StartAt > l.EndAt {
				return nil, fmt.Errorf("subtitle: start_at is greater than end_at")
			}
			lines = append(lines, l)
		} else {
			prev = line
		}
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("subtitle: file not parse correctly")
	}
	return lines, nil
}

func (s *srt) ReadFromBytes(file []byte) ([]Line, error) {
	return s.ReadFromFile(bytes.NewReader(file))
}

func newSrt() Subtitler {
	return &srt{}
}
