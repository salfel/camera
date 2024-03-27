package main

import (
	"camera-server/services/database"
	"fmt"
)

func main() {

	db := database.GetDB()

	err := db.SetupJoinTable(&database.User{}, "Streams", &database.UserStream{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(database.Session{}, database.User{}, database.Stream{})
	if err != nil {
		fmt.Println(err)
	}
}
