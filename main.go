package main

import (
	"github.com/jalawz/task-api/database"
	"github.com/jalawz/task-api/routes"
)

func main() {
	db := database.InitDB()
	routes.HandleRequests(db)
}
