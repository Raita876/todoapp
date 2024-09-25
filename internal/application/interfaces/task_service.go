package interfaces

import "github.com/raita876/todoapp/internal/application/query"

type TaskService interface {
	FindAllTasks() (*query.TaskQueryListResult, error)
}
