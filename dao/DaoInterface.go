package dao

import (
	"rest_servers/models"
	"time"
)

type TaskServerDao interface {
	CreateTask(text string, tags []string, due time.Time) int
	GetTask(id int) (models.Task, error)
	DeleteTask(id int) error
	DeleteAllTasks() error
	GetAllTasks() []models.Task
	GetTasksByTag(tag string) []models.Task
	GetTasksByDueDate(year int, month time.Month, day int) []models.Task
}
