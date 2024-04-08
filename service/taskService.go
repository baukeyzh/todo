// service/task_service.go

package service

import (
	"context"
	"github.com/baukeyzh/todo/models"
	"github.com/baukeyzh/todo/repository"
)

type TaskService interface {
	ReadTask(ctx context.Context, id string) (models.Task, error)
	ReadTasks(ctx context.Context, status string) ([]models.Task, error)
	CreateTask(ctx context.Context, task models.Task) (string, error)
	UpdateTask(ctx context.Context, task models.TaskForm) error
	DeleteTask(ctx context.Context, id string) error
	MarkTaskDone(ctx context.Context, id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) ReadTasks(ctx context.Context, status string) ([]models.Task, error) {
	if status == "done" {
		return s.repo.SelectDoneTasks(ctx)
	}
	return s.repo.SelectActiveTasks(ctx)
}

func (s *taskService) ReadTask(ctx context.Context, id string) (models.Task, error) {
	return s.repo.SelectTask(ctx, id)
}

func (s *taskService) CreateTask(ctx context.Context, task models.Task) (string, error) {
	return s.repo.InsertTask(ctx, task)
}

func (s *taskService) UpdateTask(ctx context.Context, taskForm models.TaskForm) error {
	return s.repo.UpdateTask(ctx, taskForm)
}

func (s *taskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.SetTaskDeleted(ctx, id)
}

func (s *taskService) MarkTaskDone(ctx context.Context, id string) error {
	return s.repo.SetTaskDone(ctx, id)
}
