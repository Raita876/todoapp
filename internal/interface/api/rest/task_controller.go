package rest

import (
	"net/http"

	"github.com/raita876/todoapp/internal/application/interfaces"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	service interfaces.TaskService
}

func NewTaskController(e *gin.Engine, service interfaces.TaskService) *TaskController {
	c := &TaskController{
		service: service,
	}

	e.GET("/api/v1/task", c.CreateTaskController)

	return c
}

func (tc *TaskController) CreateTaskController(c *gin.Context) {

	// TODO: ToTaskResponse 実装
	response := gin.H{
		"message": "This is task api.",
	}

	c.JSON(http.StatusOK, response)
}
