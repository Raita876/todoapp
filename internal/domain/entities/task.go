package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status struct {
	Id   int
	Name string
}

type Task struct {
	Id          uuid.UUID
	Name        string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) Validate() error {
	if t.Name == "" {
		return errors.New("name must not be empty")
	}
	if t.Description == "" {
		return errors.New("description must not be empty")
	}
	if t.CreatedAt.After(t.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}
	if t.Status.Name == "" {
		return errors.New("status_id must not be empty")
	}

	return nil
}

func (t *Task) UpdateName(name string) error {
	t.Name = name
	t.UpdatedAt = time.Now()

	return t.Validate()
}

func (t *Task) UpdateDescription(description string) error {
	t.Description = description
	t.UpdatedAt = time.Now()

	return t.Validate()
}

func (t *Task) UpdateStatus(status Status) error {
	t.Status = status
	t.UpdatedAt = time.Now()

	return t.Validate()
}

func EqualTask(src, dst *Task) bool {
	if src == nil && dst == nil {
		return true
	}

	if src == nil || dst == nil {
		return false
	}

	if src.Name != dst.Name ||
		src.Description != dst.Description ||
		src.Status != dst.Status {
		return false
	}

	return true
}
