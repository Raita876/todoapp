package repositories

import (
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/domain/entities"
)

type TaskRepository interface {
	Create(task *entities.Task) (*entities.Task, error)
	FindAll() ([]*entities.Task, error)
	FindTaskById(id uuid.UUID) (*entities.Task, error)
	Update(task *entities.Task) (*entities.Task, error)
	Delete(id uuid.UUID) error
}
