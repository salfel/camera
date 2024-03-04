package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

type User struct {
    gorm.Model
    username string
    password string
}

func GetDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return db
}

func main() {
    db := GetDB()

    db.AutoMigrate(&User{})
}
