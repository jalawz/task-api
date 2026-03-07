package database

import (
	"github.com/jalawz/task-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	database, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to the database!")
	}

	database.AutoMigrate(&models.Task{})

	DB = database
}
