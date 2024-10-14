package services

import (
	"errors"

	"github.com/google/uuid"
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

func (ts *TaskService) FindAllTasks(tasksCommand *command.FindAllTasksCommand) (*query.TaskQueryListResult, error) {
	tasks, err := ts.taskRepository.FindAll(
		tasksCommand.ContainsForName,
		tasksCommand.FilterStatusId,
		tasksCommand.SortBy,
		tasksCommand.OrderIsAsc,
	)
	if err != nil {
		return nil, err
	}

	var taskQueryListResult query.TaskQueryListResult
	for _, task := range tasks {
		taskQueryListResult.Result = append(taskQueryListResult.Result, mapper.NewTaskResultFromEntity(task))
	}

	return &taskQueryListResult, nil
}

func (ts *TaskService) FindTaskById(id uuid.UUID) (*query.TaskQueryResult, error) {
	task, err := ts.taskRepository.FindTaskById(id)
	if err != nil {
		return nil, err
	}

	var queryResult query.TaskQueryResult
	queryResult.Result = mapper.NewTaskResultFromEntity(task)

	return &queryResult, nil
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

func (ts *TaskService) UpdateTask(updateCommand *command.UpdateTaskCommand) (*command.UpdateTaskCommandResult, error) {
	task, err := ts.taskRepository.FindTaskById(updateCommand.ID)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, errors.New("task not found")
	}

	if err := task.UpdateName(updateCommand.Name); err != nil {
		return nil, err
	}

	if err := task.UpdateDescription(updateCommand.Description); err != nil {
		return nil, err
	}

	if err := task.UpdateStatus(updateCommand.StatusId); err != nil {
		return nil, err
	}

	err = task.Validate()
	if err != nil {
		return nil, err
	}

	_, err = ts.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	result := command.UpdateTaskCommandResult{
		Result: mapper.NewTaskResultFromEntity(task),
	}

	return &result, nil
}

func (ts *TaskService) DeleteTask(id uuid.UUID) error {
	return ts.taskRepository.Delete(id)
}
