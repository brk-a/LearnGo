package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task with an ID, title, body and completion status.
type Task struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`            // Unique identifier for the task.
	Title      string             `json:"title" validate:"required"` // Title of the task.
	Body       string             `json:"body" validate:"required"`  // Description of the task.
	Completed  bool               `json:"completed"`                 // Indicates whether the task is completed.
	Task_id    string             `json:"task_id"`                   //
	Created_at time.Time          `json:"created_at"`                // Date and time when the task was created.
	Updated_at time.Time          `json:"updated_at"`                // Date and time when the task was last updated.
}
