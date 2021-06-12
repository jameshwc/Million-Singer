package mysql

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/repo"
)

type mysqlSongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) repo.SongRepo {
	return &mysqlSongRepository{db: db}
}

func (m *mysqlSongRepository) CheckManyExist(songsID []int) (int64, error) {
	var count int64
	stmt := "SELECT COUNT(*) FROM songs WHERE id IN ("
	for i := range songsID {
		stmt += strconv.Itoa(songsID[i]) + ","
	}
	stmt = stmt[:len(stmt)-1] + ")"
	row := m.db.QueryRow(stmt)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (m *mysqlSongRepository) QueryByVideoID(videoID string) (id int64, err error) {
	if err = m.db.QueryRow("SELECT id FROM songs WHERE songs.video_id = ?", videoID).Scan(&id); err != nil {
		return 0, err
	}
	return id, err
}

// TODO: Fix foreign key or abandon the idea of delete data
func (m *mysqlSongRepository) Del(id int) (collectsID []int, err error) {

	tx, err := m.db.Begin()
	if err != nil {
		return
	}
	rows, err := m.db.Query("SELECT collect_id FROM collect_songs WHERE song_id = ?", id)
	if err != nil {
		tx.Rollback()
		return
	}
	defer rows.Close()

	for rows.Next() {
		collectID := 0
		rows.Scan(&collectID)
		collectsID = append(collectsID, collectID)
	}
	if len(collectsID) > 0 {
		tx.Rollback()
		return
	}

	_, err = tx.Exec("DELETE FROM songs WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (m *mysqlSongRepository) Add(videoID, name, singer, genre, language, missLyrics, startTime, endTime string, lyrics []model.Lyric) (int, error) {
	cur := time.Now()
	tx, err := m.db.Begin()
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

func (m *mysqlSongRepository) Get(songID int, hasLyrics bool) (*model.Song, error) {
	var song model.Song
	if err := m.db.QueryRow("SELECT id, video_id, name, singer, genre, language, miss_lyrics FROM songs WHERE id = ? AND deleted_at IS NULL", songID).
		Scan(&song.ID, &song.VideoID, &song.Name, &song.Singer, &song.Genre, &song.Language, &song.MissLyrics); err != nil {
		return nil, err
	} else if song.ID != songID {
		return nil, errors.New("scan id and param id are not matched")
	}
	if hasLyrics {
		rows, err := m.db.Query("SELECT `index`, line, start_at, end_at FROM lyrics WHERE song_id = ?", songID)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var lyric model.Lyric
			if err := rows.Scan(&lyric.Index, &lyric.Line, &lyric.StartAt, &lyric.EndAt); err != nil {
				return nil, err
			}
			song.Lyrics = append(song.Lyrics, &lyric)
		}
	}
	return &song, nil
}

func (m *mysqlSongRepository) Gets() (songs []*model.Song, err error) {
	rows, err := m.db.Query("SELECT id, video_id, name, singer, genre, language FROM songs WHERE deleted_at IS NULL")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var s model.Song
		if err = rows.Scan(&s.ID, &s.VideoID, &s.Name, &s.Singer, &s.Genre, &s.Language); err != nil {
			return
		}
		songs = append(songs, &s)
	}
	return
}
