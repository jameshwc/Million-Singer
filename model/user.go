package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Active    bool
	Token     string
	LastLogin *time.Time
}

func (u *User) Commit() error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func CheckDuplicateWithName(name string) bool {
	var u User
	if db.Where("name = ?", name).First(&u).Error != nil {
		return true
	}
	return false
}

func CheckDuplicateWithEmail(email string) bool {
	var u User
	if db.Where("email = ?", email).First(&u).Error != nil {
		return true
	}
	return false
}