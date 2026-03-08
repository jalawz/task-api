package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jalawz/task-api/controllers"
	"github.com/jalawz/task-api/services"
	"gorm.io/gorm"
)

func HandleRequests(db *gorm.DB) {
	r := gin.Default()

	taskService := services.NewTaskService(db)
	taskController := &controllers.TaskController{Service: taskService}

	r.GET("/tasks", taskController.ListTasks)
	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks/:id", taskController.GetTaskByID)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	r.Run(":8000")
}
