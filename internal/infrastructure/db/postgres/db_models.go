package postgre

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string
	Description string
	StatusId    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
