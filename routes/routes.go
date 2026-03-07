package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jalawz/task-api/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/tasks", controllers.ListTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.Run()
}
