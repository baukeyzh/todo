package models

import (
	"time"
)

type Task struct {
	ID       string    `bson:"_id,omitempty"`
	Title    string    `bson:"title"`
	ActiveAt time.Time `bson:"active_at"`
}
