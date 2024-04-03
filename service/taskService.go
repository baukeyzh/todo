// service/task_service.go

package service

import (
	"context"
	"github.com/baukeyzh/todo/models"
	"github.com/baukeyzh/todo/repository"
	"time"
)

type TaskService interface {
	CreateTask(ctx context.Context, task models.Task) (string, error)
	UpdateTask(ctx context.Context, id string, task models.Task) error
	DeleteTask(ctx context.Context, id string) error
	MarkTaskDone(ctx context.Context, id string) error
	GetTasksByStatus(ctx context.Context, status string, day time.Time) ([]models.Task, error)
	GetTasks(ctx context.Context, day time.Time) ([]models.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

// Здесь реализация методов интерфейса, например:

func (s *taskService) CreateTask(ctx context.Context, task models.Task) (string, error) {
	return "", nil
}

func (s *taskService) UpdateTask(ctx context.Context, id string, task models.Task) error {
	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return nil
}

func (s *taskService) MarkTaskDone(ctx context.Context, id string) error {
	return nil
}

func (s *taskService) GetTasksByStatus(ctx context.Context, status string, day time.Time) ([]models.Task, error) {
	return nil, nil
}

func (s *taskService) GetTasks(ctx context.Context, day time.Time) ([]models.Task, error) {
	tasks, err := s.repo.FindAll(ctx, day)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Остальные методы...
