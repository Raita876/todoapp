package interfaces

import (
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/application/command"
	"github.com/raita876/todoapp/internal/application/query"
)

type TaskService interface {
	CreateTask(taskCommand *command.CreateTaskCommand) (*command.CreateTaskCommandResult, error)
	FindAllTasks(findAllCommand *command.FindAllTasksCommand) (*query.TaskQueryListResult, error)
	FindTaskById(id uuid.UUID) (*query.TaskQueryResult, error)
	UpdateTask(updateCommand *command.UpdateTaskCommand) (*command.UpdateTaskCommandResult, error)
	DeleteTask(id uuid.UUID) error
}
