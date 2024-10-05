package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/interface/api/rest/dto/mapper"
	"github.com/raita876/todoapp/internal/interface/api/rest/dto/request"

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
	e.GET("/api/v1/tasks/:id", c.GetTaskByIdController)
	e.POST("/api/v1/tasks", c.CreateTaskController)

	return c
}

// @BasePath /api/v1

// PingExample godoc
// @Summary Get tasks
// @Schemes http
// @Description get tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} response.ListTaskResponse
// @Router /tasks [get]
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

// TODO: swagger 記載
func (tc *TaskController) GetTaskByIdController(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid task Id format",
		})
		return
	}

	task, err := tc.service.FindTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch task",
		})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"error": "task not found",
		})
		return
	}

	response := mapper.ToTaskResponse(task.Result)

	c.JSON(http.StatusOK, response)
}

// TODO: swagger 記載
func (tc *TaskController) CreateTaskController(c *gin.Context) {
	var createTaskRequest request.CreateTaskRequest

	if err := c.Bind(&createTaskRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
		return
	}

	taskCommand, err := createTaskRequest.ToCreateTaskCommand()
	// 現時点で到達しない処理
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid Id format",
		})
		return
	}

	result, err := tc.service.CreateTask(taskCommand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create task",
		})
		return
	}

	response := mapper.ToTaskResponse(result.Result)

	c.JSON(http.StatusCreated, response)
}

// TODO
func (tc *TaskController) PutTaskController(c *gin.Context) {

}

// TODO
func (tc *TaskController) DeleteTaskController(c *gin.Context) {

}
