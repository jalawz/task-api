package services_test

import (
	"context"
	"testing"

	"github.com/jalawz/task-api/models"
	"github.com/jalawz/task-api/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Cria um banco SQLite em memória para testes isolados
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Task{})
	return db
}

func TestTaskService_Create(t *testing.T) {
	db := setupTestDB()
	service := services.NewTaskService(db)
	ctx := context.Background()

	status := false
	task := &models.Task{
		Title:       "Test Task",
		Description: "Testing the service layer",
		Status:      &status,
	}

	createdTask, err := service.Create(ctx, task)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, createdTask.ID)
	assert.Equal(t, "Test Task", createdTask.Title)
}

func TestTaskService_GetAll(t *testing.T) {
	db := setupTestDB()
	service := services.NewTaskService(db)
	ctx := context.Background()

	// Seed data
	s1, s2 := false, true
	db.Create(&models.Task{Title: "Task 1", Status: &s1})
	db.Create(&models.Task{Title: "Task 2", Status: &s2})

	tasks, err := service.GetAll(ctx)

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestTaskService_GetByID(t *testing.T) {
	db := setupTestDB()
	service := services.NewTaskService(db)
	ctx := context.Background()

	s := false
	task := models.Task{Title: "Target Task", Status: &s}
	db.Create(&task)

	// Test found
	foundTask, err := service.GetByID(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, "Target Task", foundTask.Title)

	// Test not found
	_, err = service.GetByID(ctx, "99")
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestTaskService_Update(t *testing.T) {
	db := setupTestDB()
	service := services.NewTaskService(db)
	ctx := context.Background()

	s := false
	task := models.Task{Title: "Old Title", Status: &s}
	db.Create(&task)

	newStatus := true
	updateData := &models.Task{
		Title:  "New Title",
		Status: &newStatus,
	}

	updatedTask, err := service.Update(ctx, "1", updateData)

	assert.NoError(t, err)
	assert.Equal(t, "New Title", updatedTask.Title)
	assert.Equal(t, true, *updatedTask.Status)
}

func TestTaskService_Delete(t *testing.T) {
	db := setupTestDB()
	service := services.NewTaskService(db)
	ctx := context.Background()

	s := false
	db.Create(&models.Task{Title: "To be deleted", Status: &s})

	err := service.Delete(ctx, "1")
	assert.NoError(t, err)

	// Verify it's gone
	var task models.Task
	result := db.First(&task, 1)
	assert.Error(t, result.Error)
}
