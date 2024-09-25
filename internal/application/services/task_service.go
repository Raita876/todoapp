package services

import (
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/application/mapper"
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

func (ts *TaskService) FindAllTasks() (*query.TaskQueryListResult, error) {
	tasks, err := ts.taskRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var taskQueryListResult query.TaskQueryListResult
	for _, task := range tasks {
		taskQueryListResult.Result = append(taskQueryListResult.Result, mapper.NewTaskResultFromEntity(task))
	}

	return &taskQueryListResult, nil
}
