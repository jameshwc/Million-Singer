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

func parseSrtDuration(s string) (time.Duration, error) {
	return parseDuration(s, ",", 3)
}

func ReadSrtFromFile(i io.Reader) ([]Line, error) {
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
			var line Line
			var err error
			line.Text = scanner.Text()
			line.Index, err = strconv.Atoi(prev)
			if err != nil {
				return nil, err
			}
			boundaries := strings.Split(line.Text, srtTimeSep)
			if line.StartAt, err = parseSrtDuration(boundaries[0]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[0], err)
			}
			if line.EndAt, err = parseSrtDuration(boundaries[1]); err != nil {
				return nil, fmt.Errorf("subtitle: parsing srt duration %s failed: %w", boundaries[1], err)
			}
			if line.StartAt > line.EndAt {
				return nil, fmt.Errorf("subtitle: start_at is greater than end_at")
			}
			lines = append(lines, line)
		} else {
			prev = line
		}
	}
	if len(lines) == 0 {
		return nil, fmt.Errorf("subtitle: file not parse correctly")
	}
	return lines, nil
}

func ReadSrtFromBytes(file []byte) ([]Line, error) {
	return ReadSrtFromFile(bytes.NewReader(file))
}
