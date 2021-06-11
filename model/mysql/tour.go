package mysql

import (
	"database/sql"
	"errors"

	"strconv"
	"time"

	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/repo"
)

type mysqlTourRepository struct {
	db *sql.DB
}

func NewTourRepository(db *sql.DB) repo.TourRepo {
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

func (m *mysqlTourRepository) Gets() (tours []*model.Tour, err error) {
	rows, err := m.db.Query(`SELECT tours.title, tours.id FROM tours WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tid int
		var title string
		if err := rows.Scan(&tid, &title); err != nil {
			return nil, err
		}
		tours = append(tours, &model.Tour{ID: tid, Title: title, Collects: nil})
	}
	return
}

func (m *mysqlTourRepository) Add(collectsID []int, title string) (int, error) {
	cur := time.Now()
	tx, err := m.db.Begin()
	if err != nil {
		return 0, err
	}

	result, err := tx.Exec("INSERT INTO tours (created_at, updated_at, deleted_at, title) VALUES (?, ?, ?, ?)", cur, cur, sql.NullTime{}, title)
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

func (m *mysqlTourRepository) Del(id int) error {
	// if tour, err := m.Get(id); tour == nil {
	// return errors.New("tour has been deleted or tour id not found in database or database error" + err.Error())
	// }

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM tour_collects WHERE tour_id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("DELETE FROM tours WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
