package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID
	Name        string
	Description string
	StatusId    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
