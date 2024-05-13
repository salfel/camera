package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Streams  []Stream `gorm:"many2many:user_streams;"`
}

type Session struct {
	gorm.Model
	UserID uint
	User   User
}

type Stream struct {
	gorm.Model
	Channel      string `gorm:"unique"`
	XOrientation int    `gorm:"default:0"`
	YOrientation int    `gorm:"default:0"`
	AuthToken    string
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
