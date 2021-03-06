package cache

import (
	"strconv"

	"github.com/jameshwc/Million-Singer/pkg/constant"
)

func GetSongKey(id int, hasLyrics bool) string {
	lyrics := ""
	if hasLyrics {
		lyrics = "_LYRICS"
	}
	return constant.CACHE_SONG + "_" + strconv.Itoa(id) + lyrics
}

func GetLyricKey(id int) string {
	return constant.CACHE_LYRIC + "_" + strconv.Itoa(id)
}

func GetTourKey(id int) string {
	return constant.CACHE_TOUR + "_" + strconv.Itoa(id)
}

func GetCollectKey(id int) string {
	return constant.CACHE_COLLECT + "_" + strconv.Itoa(id)
}
