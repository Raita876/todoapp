package request

import "github.com/raita876/todoapp/internal/application/command"

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StatusId    int    `json:"status_id"`
}

func (req *CreateTaskRequest) ToCreateTaskCommand() (*command.CreateTaskCommand, error) {
	return &command.CreateTaskCommand{
		Name:        req.Name,
		Description: req.Description,
		StatusId:    req.StatusId,
	}, nil
}
