package repositories

import "github.com/raita876/todoapp/internal/domain/entities"

type TaskRepository interface {
	FindAll() ([]*entities.Task, error)
}
