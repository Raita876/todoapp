package services

import (
	"github.com/raita876/todoapp/internal/application/command"
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/application/mapper"
	"github.com/raita876/todoapp/internal/application/query"
	"github.com/raita876/todoapp/internal/domain/entities"
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

func (ts *TaskService) CreateTask(taskCommand *command.CreateTaskCommand) (*command.CreateTaskCommandResult, error) {
	task := entities.NewTask(taskCommand.Name, taskCommand.Description, taskCommand.StatusId)
	err := task.Validate()
	if err != nil {
		return nil, err
	}

	_, err = ts.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	result := command.CreateTaskCommandResult{
		Result: mapper.NewTaskResultFromEntity(task),
	}

	return &result, nil
}
