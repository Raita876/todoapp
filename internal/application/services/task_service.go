package services

import (
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/application/query"
	"github.com/raita876/todoapp/internal/domain/repositories"
)

type TaskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(taskRepository repositories.TaskRepository) interfaces.TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (ts *TaskService) FindAllTasks() (*query.TaskQueryResult, error) {
	// TODO
	return nil, nil
}
