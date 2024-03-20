package main

import (
	"camera-server/services/database"
	"fmt"
)

func main() {

	db := database.GetDB()

	err := db.AutoMigrate(database.Session{}, database.User{}, database.Visit{})
	if err != nil {
		fmt.Println(err)
	}
}
