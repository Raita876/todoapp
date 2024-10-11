package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Task struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	StatusId    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Task) TableName() string {
	return "tasks"
}
