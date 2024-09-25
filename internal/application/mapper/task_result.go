package mapper

import (
	"github.com/raita876/todoapp/internal/application/common"
	"github.com/raita876/todoapp/internal/domain/entities"
)

func NewTaskResultFromEntity(task *entities.Task) *common.TaskResult {
	if task == nil {
		return nil
	}

	return &common.TaskResult{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		StatusId:    task.StatusId,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
