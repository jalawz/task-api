package main

import (
	"github.com/jalawz/task-api/database"
	"github.com/jalawz/task-api/routes"
)

func main() {
	database.InitDB()
	routes.HandleRequests()
}
