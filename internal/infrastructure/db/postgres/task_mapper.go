package postgres

import "github.com/raita876/todoapp/internal/domain/entities"

func toDBTask(task *entities.Task) *Task {
	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		StatusId:    task.Status.Id,
		Status: Status{
			Id:   task.Status.Id,
			Name: task.Status.Name,
		},
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

func fromDBTask(dbTask *Task) *entities.Task {
	return &entities.Task{
		Id:          dbTask.Id,
		Name:        dbTask.Name,
		Description: dbTask.Description,
		Status: entities.Status{
			Id:   dbTask.Status.Id,
			Name: dbTask.Status.Name,
		},
		CreatedAt: dbTask.CreatedAt,
		UpdatedAt: dbTask.UpdatedAt,
	}
}
