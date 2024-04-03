package repository

import (
	"context"
	"github.com/baukeyzh/todo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TaskRepository interface {
	Create(ctx context.Context, task models.Task) (string, error)
	Update(ctx context.Context, id string, task models.Task) error
	Delete(ctx context.Context, id string) error
	MarkDone(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*models.Task, error)
	FindAll(ctx context.Context, day time.Time) ([]models.Task, error)
}

type mongoTaskRepository struct {
	db *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) TaskRepository {
	return &mongoTaskRepository{
		db: db.Collection("tasks"),
	}
}

// Здесь реализация методов интерфейса, например:

func (r *mongoTaskRepository) Create(ctx context.Context, task models.Task) (string, error) {
	return "", nil
}

func (r *mongoTaskRepository) Update(ctx context.Context, id string, task models.Task) error {
	return nil
}

func (r *mongoTaskRepository) MarkDone(ctx context.Context, id string) error {
	return nil
}

func (r *mongoTaskRepository) FindByID(ctx context.Context, id string) (*models.Task, error) {
	return nil, nil
}

func (r *mongoTaskRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *mongoTaskRepository) FindAll(ctx context.Context, day time.Time) ([]models.Task, error) {
	return []models.Task{
		{ID: "1", Title: "Задача 1"},
		{ID: "2", Title: "Задача 2"},
	}, nil
}

// Остальные методы...
