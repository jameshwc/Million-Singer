package model

import (
	"database/sql"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jameshwc/Million-Singer/pkg/log"
)

type Song struct {
	ID         int      `json:"id"`
	Lyrics     []*Lyric `json:"lyrics,omitempty"`
	VideoID    string   `json:"video_id"`
	StartTime  string   `json:"start_time,omitempty"`
	EndTime    string   `json:"end_time,omitempty"`
	Language   string   `json:"language"`
	Name       string   `json:"name"`
	Singer     string   `json:"singer"`
	Genre      string   `json:"genre"`
	MissLyrics string   `json:"miss_lyrics,omitempty"` // IDs (integers) with comma seperated
}

func AddSong(videoID, name, singer, genre, language, missLyrics, startTime, endTime string, lyrics []Lyric) (int, error) {
	cur := time.Now()
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec(`INSERT INTO songs (
	video_id, name, singer, genre, language, miss_lyrics, start_time, end_time, created_at, updated_at, deleted_at
	) VALUES (?,?,?,?,?,?,?,?,?,?,?)`, videoID, name, singer, genre, language, missLyrics, startTime, endTime, cur, cur, sql.NullTime{})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	stmt := "INSERT INTO lyrics (created_at, updated_at, deleted_at, `index`, line, start_at, end_at, song_id) VALUES "
	log.Info("song id ", id)
	songID := strconv.Itoa(int(id))
	curString := cur.Format("2006-01-02 15:04:05")
	for i := range lyrics {
		stmt += "('" + curString + "','" + curString + "',NULL," + strconv.Itoa(lyrics[i].Index) + ",'" + escape(lyrics[i].Line) + "'," + strconv.FormatInt(lyrics[i].StartAt.Milliseconds(), 10) + "," + strconv.FormatInt(lyrics[i].EndAt.Milliseconds(), 10) + "," + songID + "),"
	}
	stmt = stmt[:len(stmt)-1]
	result, err = tx.Exec(stmt)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (s *Song) RandomGetMissLyricID() int {
	numsStrs := strings.Split(s.MissLyrics, ",")
	// Let's assume MissLyrics had been validated so we can skip error check
	randID, _ := strconv.Atoi(numsStrs[rand.Intn(len(numsStrs))])
	return randID
}

func GetSong(songID int, hasLyrics bool) (*Song, error) {
	var song Song
	if err := db.QueryRow("SELECT id, video_id, name, singer, genre, language, miss_lyrics FROM songs WHERE id = ? AND deleted_at IS NULL", songID).
		Scan(&song.ID, &song.VideoID, &song.Name, &song.Singer, &song.Genre, &song.Language, &song.MissLyrics); err != nil {
		return nil, err
	} else if song.ID != songID {
		return nil, errors.New("scan id and param id are not matched")
	}
	if hasLyrics {
		rows, err := db.Query("SELECT `index`, line, start_at, end_at FROM lyrics WHERE song_id = ?", songID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var lyric Lyric
			if err := rows.Scan(&lyric.Index, &lyric.Line, &lyric.StartAt, &lyric.EndAt); err != nil {
				log.Info(err)
				return nil, err
			}
			song.Lyrics = append(song.Lyrics, &lyric)
		}
	}
	return &song, nil
}

func CheckSongsExist(songsID []int) (int64, error) {
	var count int64
	stmt := "SELECT COUNT(*) FROM songs WHERE id IN ("
	for i := range songsID {
		stmt += strconv.Itoa(songsID[i]) + ","
	}
	stmt = stmt[:len(stmt)-1] + ")"
	row := db.QueryRow(stmt)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func QuerySongByVideoID(videoID string) (id int64, err error) {
	if err = db.QueryRow("SELECT id FROM songs WHERE songs.video_id = ?", videoID).Scan(&id); err != nil {
		return 0, err
	}
	return id, err
}
