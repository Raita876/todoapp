package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Status struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func (Status) TableName() string {
	return "task_status"
}

type Task struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	StatusId    int
	Status      Status `gorm:"foreignKey:StatusId;references:Id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Task) TableName() string {
	return "tasks"
}
