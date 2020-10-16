package model

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Active    bool
	LastLogin *time.Time
}

func (u *User) UpdateLoginStatus() error {
	// TODO
	return nil
}
