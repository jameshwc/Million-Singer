package subtitle

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/jameshwc/Million-Singer/model"
)

const (
	srtTimeSep = " --> "
)

func parseSrtDuration(s string) (time.Duration, error) {
	return parseDuration(s, ",", 3)
}

func ReadSrtFromFile(i io.Reader) ([]model.Lyric, error) {
	scanner := bufio.NewScanner(i)
	var lyrics []model.Lyric
	var prev string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, srtTimeSep) {
			scanner.Scan()
			var lyric model.Lyric
			var err error
			lyric.Line = scanner.Text()
			lyric.Index, err = strconv.Atoi(prev)
			if err != nil {
				return nil, err
			}
			boundaries := strings.Split(line, srtTimeSep)
			if lyric.StartAt, err = parseSrtDuration(boundaries[0]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[0], err)
			}
			if lyric.EndAt, err = parseSrtDuration(boundaries[1]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[1], err)
			}
			if lyric.StartAt > lyric.EndAt {
				return nil, fmt.Errorf("subtitle: start_at is greater than end_at")
			}
			lyrics = append(lyrics, lyric)
		} else {
			prev = line
		}
	}
	if len(lyrics) == 0 {
		return nil, fmt.Errorf("subtitle: file not parse correctly")
	}
	return lyrics, nil
}

func ReadSrtFromBytes(file []byte) ([]model.Lyric, error) {
	return ReadSrtFromFile(bytes.NewReader(file))
}
