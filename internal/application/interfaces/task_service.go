package interfaces

import (
	"github.com/raita876/todoapp/internal/application/command"
	"github.com/raita876/todoapp/internal/application/query"
)

type TaskService interface {
	CreateTask(taskCommand *command.CreateTaskCommand) (*command.CreateTaskCommandResult, error)
	FindAllTasks() (*query.TaskQueryListResult, error)
}
