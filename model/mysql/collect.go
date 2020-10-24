package mysql

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/repo"
)

type mysqlCollectRepository struct {
	db *sql.DB
}

func NewCollectRepository(db *sql.DB) repo.CollectRepo {
	return &mysqlCollectRepository{db: db}
}

func (m *mysqlCollectRepository) Get(id int) (*model.Collect, error) {
	scanID := 0
	title := ""
	if err := m.db.QueryRow("SELECT id, title FROM collects WHERE id = ? AND deleted_at IS NULL", id).Scan(&scanID, &title); err != nil {
		return nil, err
	} else if scanID != id {
		return nil, errors.New("scan id and param id are not matched")
	}

	var collect model.Collect
	collect.ID = id
	collect.Title = title

	rows, err := m.db.Query(`SELECT songs.id as song_id, songs.video_id, songs.name, songs.singer, songs.language, songs.genre
				FROM songs
				INNER JOIN collect_songs ON collect_songs.collect_id = ? AND collect_songs.song_id = songs.id`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song
		rows.Scan(&song.ID, &song.VideoID, &song.Name, &song.Singer, &song.Language, &song.Genre)
		collect.Songs = append(collect.Songs, &song)
	}
	return &collect, nil
}

func (m *mysqlCollectRepository) Add(title string, songsID []int) (int, error) {
	cur := time.Now()
	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec("INSERT INTO collects (title, created_at, updated_at, deleted_at) VALUES (?,?,?,?)", title, cur, cur, sql.NullTime{})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	stmt := "INSERT INTO collect_songs (collect_id, song_id) VALUES "
	collectID := strconv.Itoa(int(id))
	for i := range songsID {
		stmt += "(" + collectID + "," + strconv.Itoa(songsID[i]) + "),"
	}
	stmt = stmt[:len(stmt)-1]

	result, err = tx.Exec(stmt)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}

func (m *mysqlCollectRepository) CheckManyExist(collectsID []int) (int64, error) {
	var count int64
	stmt := "SELECT COUNT(*) FROM collects WHERE id IN ("
	for i := range collectsID {
		stmt += strconv.Itoa(collectsID[i]) + ","
	}
	stmt = stmt[:len(stmt)-1] + ")"
	row := m.db.QueryRow(stmt)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
