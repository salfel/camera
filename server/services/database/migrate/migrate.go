package main

import (
	"fmt"

	"camera-server/services/database"
)

func main() {

    db := database.GetDB()

    db.AutoMigrate(database.Session{}, database.User{})

    fmt.Println("created")

    db.Create(&database.User{Username: "Felix", Password: "Felix"})

    var user *database.User
    db.First(&user)

    fmt.Println(user)
}
