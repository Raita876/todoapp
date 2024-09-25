package mapper

import (
	"github.com/raita876/todoapp/internal/application/common"
	"github.com/raita876/todoapp/internal/interface/api/rest/dto/response"
)

func ToTaskResponse(task *common.TaskResult) *response.TaskResponse {
	return &response.TaskResponse{
		Id:          task.Id.String(),
		Name:        task.Name,
		Description: task.Description,
		StatusId:    task.StatusId,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func ToTaskListResponse(tasks []*common.TaskResult) *response.ListTaskResponse {
	var responseList []*response.TaskResponse
	for _, task := range tasks {
		responseList = append(responseList, ToTaskResponse(task))
	}

	return &response.ListTaskResponse{Tasks: responseList}
}
