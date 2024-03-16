package main

import (
	"clean-architecture/db"
	"clean-architecture/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})

}
