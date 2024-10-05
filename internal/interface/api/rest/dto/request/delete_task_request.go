package request

import (
	"github.com/google/uuid"
)

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}
