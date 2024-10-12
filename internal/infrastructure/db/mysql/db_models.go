package mysql

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Task struct {
	ID          uuid.UUID `gorm:"column:id;type:char(36);primaryKey" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description;not null" json:"description"`
	StatusID    int32     `gorm:"column:status_id;not null" json:"status_id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
