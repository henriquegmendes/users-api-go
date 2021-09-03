package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@/go-users-api"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = database
}
