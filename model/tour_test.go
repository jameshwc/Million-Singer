package model

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	gormmock "github.com/jameshwc/Million-Singer/pkg/gorm-sqlmock"
)

var tourColumns = []string{"id", "updated_at", "created_at", "deleted_at"}
var collectColumns = append(tourColumns, "title")
var testCollectID = []driver.Value{1, 2, 3}

func setupTestDatabase(t *testing.T) sqlmock.Sqlmock {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock fail to new", err.Error())
		return nil
	}
	Setup(mockdb)
	return mock
}
func TestGetTourSuccess(t *testing.T) {

	mock := setupTestDatabase(t)
	if mock == nil {
		return
	}

	mock.ExpectQuery("SELECT (.+) FROM `tours` WHERE (.+) AND `tours`.`deleted_at` IS NULL ORDER BY `tours`.`id` LIMIT 1").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(tourColumns).AddRow(1, time.Now(), time.Now(), nil),
		)
	mock.ExpectQuery("SELECT (.+) FROM `tour_collects` WHERE `tour_collects`.`tour_id` = \\?").
		WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"tour_id", "collect_id"}).AddRow(1, 1).AddRow(1, 2).AddRow(1, 3),
	)
	mock.ExpectQuery("SELECT (.+) FROM `collects`").WithArgs(testCollectID...).WillReturnRows(sqlmock.NewRows(collectColumns).
		AddRow(1, time.Now(), time.Now(), nil, "hello").
		AddRow(2, time.Now(), time.Now(), nil, "world").
		AddRow(3, time.Now(), time.Now(), nil, "hi"))
	tour, err := GetTour(1)
	if err != nil {
		t.Error("errors happened when get tour: ", err.Error())
	}

	if tour.ID != 1 {
		t.Error("error that tour id is not matched")
	}

	for i := range tour.Collects {
		if v := testCollectID[i].(int); v != int(tour.Collects[i].ID) {
			t.Error("collect id is not matched", v, int(tour.Collects[i].ID))
		}
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("errors that expectations were not met: ", err.Error())
	}
	// web_1    | [0.282ms] [rows:3] SELECT * FROM `tour_collects` WHERE `tour_collects`.`tour_id` = 5
	// web_1    | [0.452ms] [rows:3] SELECT * FROM `collects` WHERE `collects`.`id` IN (11,12,13) AND `collects`.`deleted_at` IS NULL
	// web_1    | [1.479ms] [rows:1] SELECT * FROM `tours` WHERE id = 5 AND `tours`.`deleted_at` IS NULL ORDER BY `tours`.`id` LIMIT 1
}

func TestGetTourFailToGetLevel(t *testing.T) {

	mock := setupTestDatabase(t)
	if mock == nil {
		return
	}

	mock.ExpectQuery("SELECT (.+) FROM `tours` WHERE (.+) AND `tours`.`deleted_at` IS NULL ORDER BY `tours`.`id` LIMIT 1").
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows(tourColumns).AddRow(1, time.Now(), time.Now(), nil),
		)
	mock.ExpectQuery("SELECT (.+) FROM `tour_collects` WHERE `tour_collects`.`tour_id` = \\?").
		WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"tour_id", "collect_id"}),
	)
	// TODO: Finish it
}

func TestAddTour(t *testing.T) {

	mock := setupTestDatabase(t)
	if mock == nil {
		return
	}

	mock.ExpectExec("INSERT INTO `tours` \\(`created_at`,`updated_at`,`deleted_at`\\) VALUES").
		WithArgs(gormmock.AnyTime{}, gormmock.AnyTime{}, nil).WillReturnResult(sqlmock.NewResult(1, 1))

	if _, err := AddTour(nil); err != nil {
		t.Error("errors: tour commit: ", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("errors that expectations were not met: ", err.Error())
	}
}

func TestGetTotalTours(t *testing.T) {

	mock := setupTestDatabase(t)
	if mock == nil {
		return
	}

	mock.ExpectQuery("SELECT (.+) FROM `tours` WHERE `tours`\\.`deleted_at` IS NULL").
		WillReturnRows(sqlmock.NewRows(tourColumns).AddRow(1, time.Now(), time.Now(), nil).AddRow(2, time.Now(), time.Now(), nil))

	if v, err := GetTotalTours(); err != nil {
		t.Error("error: GetTotalTours: ", err.Error())
	} else if v != 2 {
		t.Error("error: GetTotalTours not return correct number: expected 2 and actual ", v)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error("errors that expectations were not met: ", err.Error())
	}
}
