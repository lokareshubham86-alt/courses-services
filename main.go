package main

import (
	"courses-services/database"
)

// "course-service/database"go mod

func main() {
	database.ConnectToDatabase()
	database.CreateTable()
}
