package model

import (
	"crypto/sha1"
	"fmt"
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

func CheckDuplicateUserWithName(name string) bool {
	var u User
	if db.Where("name = ?", name).First(&u).Error != nil {
		return true
	}
	return false
}

func CheckDuplicateUserWithEmail(email string) bool {
	var u User
	if db.Where("email = ?", email).First(&u).Error != nil {
		return true
	}
	return false
}

func encrypt(pw string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(pw)))
}

func AuthUser(username string, password string) (*User, error) {
	var u User
	if err := db.Where("name = ?", username).Where("password = ?", encrypt(password)).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
