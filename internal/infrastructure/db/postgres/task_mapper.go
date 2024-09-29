package postgre

import "github.com/raita876/todoapp/internal/domain/entities"

func toDBTask(task *entities.Task) *Task {
	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		StatusId:    task.StatusId,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func fromDBTask(dbTask *Task) *entities.Task {
	return &entities.Task{
		Id:          dbTask.Id,
		Name:        dbTask.Name,
		Description: dbTask.Description,
		StatusId:    dbTask.StatusId,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}
}
