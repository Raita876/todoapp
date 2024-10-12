package mysql

import (
	"github.com/raita876/todoapp/internal/domain/entities"
)

func toDBTask(task *entities.Task) *Task {
	return &Task{
		ID:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		StatusID:    int32(task.StatusId),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func fromDBTask(dbTask *Task) *entities.Task {
	return &entities.Task{
		Id:          dbTask.ID,
		Name:        dbTask.Name,
		Description: dbTask.Description,
		StatusId:    int(dbTask.StatusID),
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}
