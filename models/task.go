package models

import (
	"time"
)

type Task struct {
	ID        string    `bson:"_id,omitempty"`
	Title     string    `bson:"title" validate:"required,min=3,max=200"`
	ActiveAt  time.Time `bson:"active_at" validate:"required"`
	IsDeleted bool      `bson:"is_deleted"`
	IsDone    bool      `bson:"is_done"`
}

type TaskForm struct {
	ID        string    `bson:"_id,omitempty" validate:"required,min=3,max=30"`
	Title     string    `bson:"title" validate:"required,min=3,max=200"`
	ActiveAt  time.Time `bson:"active_at" validate:"required"`
	IsDeleted bool      `bson:"is_deleted"`
}

type IdValidate struct {
	ID string `bson:"_id" validate:"required,min=3,max=30"`
}

type StatusValidate struct {
	Status string `bson:"status" validate:"min=3,max=30"`
}
