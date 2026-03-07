package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jalawz/task-api/controllers"
	"github.com/jalawz/task-api/database"
)

func main() {
	r := gin.Default()

	database.InitDB()

	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
