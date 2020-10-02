package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Active    bool
	LastLogin *time.Time
}

func IsUserNameDuplicate(name string) bool {
	cnt := 0
	if db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = ?)", name).Scan(&cnt) == sql.ErrNoRows {
		return false
	}
	return true
}

func IsUserEmailDuplicate(email string) bool {
	cnt := 0
	if db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&cnt) == sql.ErrNoRows {
		return false
	}
	return true
}

func encrypt(pw string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(pw)))
}

func AddUser(name, email, password string) (int64, error) {
	password = encrypt(password)
	result, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", name, email, password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func AuthUser(username string, password string) (*User, error) {
	var u User
	row := db.QueryRow("SELECT id, name, last_login WHERE name = ? AND password = ?", username, encrypt(password))
	if err := row.Scan(&u.ID, &u.Name, &u.LastLogin); err != nil {
		return nil, err
	}
	return &u, nil
}

func (u *User) UpdateLoginStatus() error {
	// TODO
	return nil
}
