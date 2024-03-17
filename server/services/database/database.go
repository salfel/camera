package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

type Session struct {
	gorm.Model
	UserID uint
	User   User
}

type Visit struct {
	gorm.Model
	UserID  uint
	Channel string
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	if DB != nil {
		return DB
	}

	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
