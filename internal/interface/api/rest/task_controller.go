package rest

import (
	"net/http"

	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/interface/api/rest/dto/mapper"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service interfaces.TaskService
}

func NewTaskController(e *gin.Engine, service interfaces.TaskService) *TaskController {
	c := &TaskController{
		service: service,
	}

	e.GET("/api/v1/tasks", c.GetAllTasksController)

	return c
}

func (tc *TaskController) GetAllTasksController(c *gin.Context) {

	tasks, err := tc.service.FindAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch task",
		})
		return
	}

	response := mapper.ToTaskListResponse(tasks.Result)

	c.JSON(http.StatusOK, response)
}
