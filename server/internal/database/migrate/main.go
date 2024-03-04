package main

import (
    "camera-server/internal/database"
)

func main() {
    db := database.GetDB()

    db.AutoMigrate(&database.User{})
}
