package entities

import (
	"errors"
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

	return nil
}

func NewTask(name, description string, statusId int) *Task {
	return &Task{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		StatusId:    statusId,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
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

func (t *Task) UpdateStatus(statusId int) error {
	t.StatusId = statusId
	t.UpdatedAt = time.Now()

	return t.Validate()
}
