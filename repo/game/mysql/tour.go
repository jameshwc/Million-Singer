package game

import (
	"database/sql"
	"errors"

	"strconv"
	"time"

	"github.com/jameshwc/Million-Singer/model"
)

type mysqlTourRepository struct {
	db *sql.DB
}

func NewMySQLTourRepository(db *sql.DB) *mysqlTourRepository {
	return &mysqlTourRepository{db: db}
}

func (m *mysqlTourRepository) Get(id int) (*model.Tour, error) {

	scanID := 0
	if err := m.db.QueryRow("SELECT id FROM tours WHERE id = ? AND deleted_at IS NULL", id).Scan(&scanID); err != nil {
		return nil, err
	} else if scanID != id {
		return nil, errors.New("scan id and param id are not matched")
	}

	rows, err := m.db.Query(`SELECT collects.title, collects.id FROM collects 
								INNER JOIN tour_collects ON collects.id = tour_collects.collect_id 
								AND tour_collects.tour_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tour model.Tour
	tour.ID = id
	for rows.Next() {
		title, id := "", 0
		rows.Scan(&title, &id)
		tour.Collects = append(tour.Collects, &model.Collect{Title: title, ID: id})
	}
	return &tour, nil
}

func (m *mysqlTourRepository) GetTotal() (int, error) {
	count := 0
	err := m.db.QueryRow("SELECT COUNT(*) FROM tours WHERE deleted_at IS NULL").Scan(&count)
	return count, err
}

func (m *mysqlTourRepository) Add(collectsID []int) (int, error) {
	cur := time.Now()
	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO tours (created_at, updated_at, deleted_at) VALUES (?, ?, ?)", cur, cur, sql.NullTime{})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	stmt := "INSERT INTO tour_collects VALUES "
	tourIDstring := strconv.Itoa(int(id))
	for i := range collectsID {
		stmt += "(" + tourIDstring + "," + strconv.Itoa(collectsID[i]) + "),"
	}
	stmt = stmt[:len(stmt)-1]

	result, err = tx.Exec(stmt)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), tx.Commit()
}
