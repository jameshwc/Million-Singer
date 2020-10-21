package mysql

import (
	"crypto/sha1"
	"database/sql"
	"fmt"

	"github.com/jameshwc/Million-Singer/model"
)

type mysqlUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *mysqlUserRepository {
	return &mysqlUserRepository{db: db}
}

func (m *mysqlUserRepository) IsNameDuplicate(name string) bool {
	cnt := 0
	if m.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = ?)", name).Scan(&cnt); cnt == 0 {
		return false
	}
	return true
}

func (m *mysqlUserRepository) IsEmailDuplicate(email string) bool {
	cnt := 0
	if m.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&cnt); cnt == 0 {
		return false
	}
	return true
}

func encrypt(pw string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(pw)))
}

func (m *mysqlUserRepository) Add(name, email, password string) (int64, error) {
	password = encrypt(password)
	result, err := m.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *mysqlUserRepository) Auth(username string, password string) (*model.User, error) {
	var u model.User
	row := m.db.QueryRow("SELECT id, name, last_login FROM users WHERE name = ? AND password = ?", username, encrypt(password))
	if err := row.Scan(&u.ID, &u.Name, &u.LastLogin); err != nil {
		return nil, err
	}
	return &u, nil
}
