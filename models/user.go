package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint
	Name     string `gorm:"not null"`
	Age      int    `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password []byte `gorm:"not null"`
	Address  string `gorm:"not null"`
}

func (u User) HasInvalidPassword(password string) bool {
	error := bcrypt.CompareHashAndPassword(u.Password, []byte(password))

	return error != nil
}
