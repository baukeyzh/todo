package repository

import (
	"context"
	"errors"
	"github.com/baukeyzh/todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type TaskRepository interface {
	SelectTask(ctx context.Context, id string) (models.Task, error)
	SelectActiveTasks(ctx context.Context) ([]models.Task, error)
	SelectDoneTasks(ctx context.Context) ([]models.Task, error)
	InsertTask(ctx context.Context, task models.Task) (string, error)
	UpdateTask(ctx context.Context, taskForm models.TaskForm) error
	SetTaskDeleted(ctx context.Context, id string) error
	SetTaskDone(ctx context.Context, id string) error
}

type mongoTaskRepository struct {
	db *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) TaskRepository {
	return &mongoTaskRepository{
		db: db.Collection("tasks"),
	}
}
func (r *mongoTaskRepository) SelectActiveTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	now := time.Now()
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location())

	filter := bson.M{
		"active_at": bson.M{
			"$lte": endOfDay,
		},
	}

	opts := options.Find().SetSort(bson.D{{"active_at", 1}})

	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *mongoTaskRepository) SelectDoneTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	filter := bson.M{"isDone": true}

	opts := options.Find().SetSort(bson.D{{"active_at", 1}})

	cursor, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *mongoTaskRepository) SelectTask(ctx context.Context, id string) (models.Task, error) {
	var task models.Task

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("converting ID error: %v", err)
		return task, err
	}

	filter := bson.M{"_id": objID}
	err = r.db.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return task, err
		}
		log.Printf("search error: %v", err)
		return task, err
	}

	return task, nil
}

func (r *mongoTaskRepository) InsertTask(ctx context.Context, task models.Task) (string, error) {
	result, err := r.db.InsertOne(ctx, task)
	if err != nil {
		log.Printf("insert error: %v", err)
		return "", err
	}

	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID.Hex(), nil
}

func (r *mongoTaskRepository) UpdateTask(ctx context.Context, taskForm models.TaskForm) error {
	objID, err := primitive.ObjectIDFromHex(taskForm.ID)
	if err != nil {
		log.Printf("error converting id to ObjectID: %v", err)
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":     taskForm.Title,
			"active_at": taskForm.ActiveAt,
		},
	}

	result, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		update,
	)

	if err != nil {
		log.Printf("update error: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *mongoTaskRepository) SetTaskDeleted(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": bson.M{"is_deleted": true}}

	result, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *mongoTaskRepository) SetTaskDone(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{"$set": bson.M{"is_done": true}}

	result, err := r.db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
