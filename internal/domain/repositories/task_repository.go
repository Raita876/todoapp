package repositories

import "github.com/raita876/todoapp/internal/domain/entities"

type TaskRepository interface {
	Create(task *entities.Task) (*entities.Task, error)
	FindAll() ([]*entities.Task, error)
}
