package services

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    Username string
    Password string
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

    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    return db
}

func main() {
    db := GetDB()

    db.AutoMigrate(&User{})
    db.AutoMigrate(&Session{})
}
