package services

import (
	"context"

	"github.com/jalawz/task-api/models"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/mock_task_service.go -package=mocks github.com/jalawz/task-api/services TaskService

type TaskService interface {
	Create(ctx context.Context, task *models.Task) (models.Task, error)
	GetAll(ctx context.Context) ([]models.Task, error)
	GetByID(ctx context.Context, id string) (models.Task, error)
	Update(ctx context.Context, id string, taskUpdate *models.Task) (models.Task, error)
	Delete(ctx context.Context, id string) error
}

type taskService struct {
	DB *gorm.DB
}

func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{DB: db}
}

func (s *taskService) Create(ctx context.Context, task *models.Task) (models.Task, error) {
	err := s.DB.WithContext(ctx).Create(task).Error
	return *task, err
}

func (s *taskService) GetAll(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	err := s.DB.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (s *taskService) GetByID(ctx context.Context, id string) (models.Task, error) {
	var task models.Task
	err := s.DB.WithContext(ctx).First(&task, id).Error
	return task, err
}

func (s *taskService) Update(ctx context.Context, id string, taskUpdate *models.Task) (models.Task, error) {
	var task models.Task
	if err := s.DB.WithContext(ctx).First(&task, id).Error; err != nil {
		return models.Task{}, err
	}

	// O GORM atualizará apenas os campos não-zero no banco, 
	// mas como o Status é um ponteiro, ele poderá atualizar para false corretamente.
	if err := s.DB.WithContext(ctx).Model(&task).Updates(taskUpdate).Error; err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (s *taskService) Delete(ctx context.Context, id string) error {
	result := s.DB.WithContext(ctx).Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
