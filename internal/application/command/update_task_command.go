package command

import "github.com/google/uuid"

type UpdateTaskCommand struct {
	ID          uuid.UUID
	Name        string
	Description string
	StatusId    int
}
