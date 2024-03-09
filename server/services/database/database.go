package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    Username string
    Password string

    SessionID uint
}

type Session struct {
    gorm.Model
    User User 
}

var DB *gorm.DB = &gorm.DB{}

func GetDB() *gorm.DB {
    if *DB != (gorm.DB{}) {
        return DB
    }

    db, err := gorm.Open(sqlite.Open("../../database.sqlite"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    return db
}
