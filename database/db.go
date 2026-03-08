package database

import (
	"github.com/jalawz/task-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("Fail to connect to the database!")
	}

	db.AutoMigrate(&models.Task{})
	return db
}
