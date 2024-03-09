package main

import (
	"camera-server/services/database"
)

func main() {

    db := database.GetDB()

    db.AutoMigrate(database.Session{}, database.User{})
}
