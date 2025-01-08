package models

type Task struct {
	ID int64 `json:"_id" bson:"_id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Completed bool `json:"completed"`
}