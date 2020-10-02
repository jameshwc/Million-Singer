package model

import (
	"time"

	"gorm.io/gorm"
)

type Lyric struct {
	gorm.Model `json:"-"`
	Index      int           `json:"index"`
	Line       string        `sql:"type:VARCHAR(128) CHARACTER SET utf8mb4 COLLATE utf8_general_ci" json:"line"`
	StartAt    time.Duration `json:"start_time"`
	EndAt      time.Duration `json:"end_time"`
	SongID     uint          `json:"-"`
}

// func GetLyricsWithSongID(songID int) (lyrics []*Lyric, err error) {
// 	if err = db.Where("song_id = ?", songID).Find(&lyrics).Error; err != nil {
// 		return nil, err
// 	}
// 	return
// }

func escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}
