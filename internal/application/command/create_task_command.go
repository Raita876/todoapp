package command

import "github.com/google/uuid"

type CreateTaskCommand struct {
	Name        string
	Description string
	StatusId    int
}

// TODO: パッケージ分割
type UpdateTaskCommand struct {
	ID          uuid.UUID
	Name        string
	Description string
	StatusId    int
}
