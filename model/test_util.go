package model

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type testData struct {
	ToursID      []driver.Value
	CollectsID   []driver.Value
	SongsID      []driver.Value
	TourCollect  *sqlmock.Rows
	CollectSongs *sqlmock.Rows
}

var tourColumns = []string{"id", "updated_at", "created_at", "deleted_at"}
var collectColumns = append(tourColumns, "title")
var tourCollectColumns = []string{"tour_id", "collect_id"}
var testCollectID = []driver.Value{1, 2, 3}

var tourCollectTestData = sqlmock.NewRows(tourCollectColumns).AddRow(1, 1).AddRow(1, 2).AddRow(1, 3)

func setupTestDatabase(t *testing.T) sqlmock.Sqlmock {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("sqlmock fail to new", err.Error())
		return nil
	}
	Setup(mockdb)
	return mock
}

const (
	testGetTourCollectNum = 5
)

func setupTestCases(t *testing.T) {
	testTour := &Tour{}
	testTour.ID = 1
	testTour.Collects = make([]*Collect, testGetTourCollectNum)
	idx := uint(1)
	cur := time.Now()
	for i := range testTour.Collects {
		testTour.Collects[i] = &Collect{
			1,
			cur,
			cur,
			sql.NullTime{},
			"hello world", // title
			nil,           // Songs
		}
		idx++
	}
}

func getTestTour() *Tour {
	testTour := &Tour{}
	testTour.ID = 1
	testTour.Collects = make([]*Collect, testGetTourCollectNum)
	idx := 1
	for i := range testTour.Collects {
		testTour.Collects[i] = &Collect{
			ID:        idx,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: sql.NullTime{},
			Title:     "hello world", // title
		}
		idx++
	}
	return testTour
}

func getCollectTestCase(t *testing.T) *Collect {
	return &Collect{Title: "hello, world"}
}
