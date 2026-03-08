package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/jalawz/task-api/controllers"
	"github.com/jalawz/task-api/mocks"
	"github.com/jalawz/task-api/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetTaskByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.GET("/tasks/:id", controller.GetTaskByID)

	status := false
	task := models.Task{ID: 1, Title: "Test Task", Status: &status}
	
	mockService.EXPECT().
		GetByID(gomock.Any(), "1").
		Return(task, nil)

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Task", response.Title)
}

func TestGetTaskByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.GET("/tasks/:id", controller.GetTaskByID)

	mockService.EXPECT().
		GetByID(gomock.Any(), "99").
		Return(models.Task{}, gorm.ErrRecordNotFound)

	req, _ := http.NewRequest("GET", "/tasks/99", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.POST("/tasks", controller.CreateTask)

	status := false
	taskInput := models.Task{Title: "New Task", Status: &status}
	expectedTask := models.Task{ID: 1, Title: "New Task", Status: &status}

	mockService.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(expectedTask, nil)

	body, _ := json.Marshal(taskInput)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, uint(1), response.ID)
}

func TestListTasks_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.GET("/tasks", controller.ListTasks)

	status := true
	tasks := []models.Task{{ID: 1, Title: "Task 1", Status: &status}}
	
	mockService.EXPECT().
		GetAll(gomock.Any()).
		Return(tasks, nil)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response, 1)
}

func TestUpdateTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.PUT("/tasks/:id", controller.UpdateTask)

	status := true
	taskUpdate := models.Task{Title: "Updated Task", Status: &status}
	expectedTask := models.Task{ID: 1, Title: "Updated Task", Status: &status}

	mockService.EXPECT().
		Update(gomock.Any(), "1", gomock.Any()).
		Return(expectedTask, nil)

	body, _ := json.Marshal(taskUpdate)
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTaskService(ctrl)
	controller := &controllers.TaskController{Service: mockService}

	r := gin.Default()
	r.DELETE("/tasks/:id", controller.DeleteTask)

	mockService.EXPECT().
		Delete(gomock.Any(), "1").
		Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
