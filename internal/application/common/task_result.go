package common

import (
	"time"

	"github.com/google/uuid"
)

type TaskResult struct {
	Id          uuid.UUID
	Name        string
	Description string
	StatusId    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
