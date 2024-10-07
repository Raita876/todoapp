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

func NewTaskController(r *gin.Engine, service interfaces.TaskService) *TaskController {
	c := &TaskController{
		service: service,
	}

	r.GET("/api/v1/tasks", c.GetAllTasksController)
	r.GET("/api/v1/tasks/:id", c.GetTaskByIdController)
	r.POST("/api/v1/tasks", c.CreateTaskController)
	r.PUT("/api/v1/tasks", c.PutTaskController)
	r.DELETE("/api/v1/tasks", c.DeleteTaskController)

	return c
}

// @BasePath /api/v1

// @Summary Get tasks
// @Schemes http
// @Description get tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} response.ListTaskResponse
// @Failure 500 {object} map[string]string
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

// @Summary Get task by id
// @Schemes http
// @Description Get task by id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "task id"
// @Success 200 {object} response.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks/{id} [get]
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

// @Summary Create task
// @Schemes http
// @Description Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body request.CreateTaskRequest true "request body"
// @Success 201 {object} response.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [post]
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

// @Summary Update task
// @Schemes http
// @Description Update task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body request.UpdateTaskRequest true "request body"
// @Success 200 {object} response.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /tasks [put]
func (tc *TaskController) PutTaskController(c *gin.Context) {
	var updateTaskRequest request.UpdateTaskRequest

	if err := c.Bind(&updateTaskRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
		return
	}

	updateTaskCommand, err := updateTaskRequest.ToUpdateTaskCommand()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid task Id format",
		})
		return
	}

	commandResult, err := tc.service.UpdateTask(updateTaskCommand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update task",
		})
		return
	}

	response := mapper.ToTaskResponse(commandResult.Result)

	c.JSON(http.StatusOK, response)
}

// @Summary Delete task
// @Schemes http
// @Description Delete task
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body request.DeleteTaskRequest true "request body"
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /tasks [delete]
func (tc *TaskController) DeleteTaskController(c *gin.Context) {
	var deleteTaskRequest request.DeleteTaskRequest

	if err := c.Bind(&deleteTaskRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
		return
	}

	err := tc.service.DeleteTask(deleteTaskRequest.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete task",
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
