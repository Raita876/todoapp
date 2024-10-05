package request

import (
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/application/command"
)

type UpdateTaskRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StatusId    int       `json:"status_id"`
}

func (req *UpdateTaskRequest) ToUpdateTaskCommand() (*command.UpdateTaskCommand, error) {
	return &command.UpdateTaskCommand{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		StatusId:    req.StatusId,
	}, nil
}
