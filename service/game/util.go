package game

import (
	"strconv"
	"strings"
)

func checkDuplicateInts(nums []int) bool {
	c := make(map[int]bool)
	for _, n := range nums {
		if _, ok := c[n]; ok {
			return true
		}
		c[n] = true
	}
	return false
}

func lyricsJoin(lyrics []int) string {
	s := make([]string, len(lyrics))
	for i, v := range lyrics {
		s[i] = strconv.Itoa(v)
	}
	return strings.Join(s, ",")
}

func findMax(l []int) int {
	var max int
	for i := range l {
		if l[i] > max {
			max = l[i]
		}
		if l[i] < 0 {
			return l[i]
		}
	}
	return max
}
