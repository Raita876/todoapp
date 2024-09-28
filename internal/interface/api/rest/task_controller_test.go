package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/application/mapper"
	"github.com/raita876/todoapp/internal/application/query"
	"github.com/raita876/todoapp/internal/domain/entities"
	"github.com/raita876/todoapp/internal/interface/api/rest/dto/response"
	"github.com/stretchr/testify/mock"
)

var (
	now = time.Now()
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) FindAllTasks() (*query.TaskQueryListResult, error) {
	args := m.Called()

	taskQueryListResult := &query.TaskQueryListResult{}

	for _, task := range args.Get(0).([]*entities.Task) {
		taskQueryListResult.Result = append(taskQueryListResult.Result, mapper.NewTaskResultFromEntity(task))
	}

	return taskQueryListResult, args.Error(1)
}

func TestNewTaskController(t *testing.T) {
	mockTaskService := &MockTaskService{}

	type args struct {
		e       *gin.Engine
		service interfaces.TaskService
	}
	tests := []struct {
		name string
		args args
		want *TaskController
	}{
		{
			name: "nomal",
			args: args{
				e:       gin.Default(),
				service: mockTaskService,
			},
			want: &TaskController{
				service: mockTaskService,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskController(tt.args.e, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskController_GetAllTasksController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	rec := httptest.NewRecorder()
	c, e := gin.CreateTestContext(rec)
	c.Request = req

	mockService := new(MockTaskService)

	expectedTasks := []*entities.Task{
		{
			Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
			Name:        "Task One",
			Description: "This is the first task",
			StatusId:    1,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	ctrl := NewTaskController(e, mockService)
	mockService.On("FindAllTasks").Return(expectedTasks, nil)

	var expectedListResponse response.ListTaskResponse
	for _, task := range expectedTasks {
		expectedListResponse.Tasks = append(expectedListResponse.Tasks,
			&response.TaskResponse{
				Id:          task.Id.String(),
				Name:        task.Name,
				Description: task.Description,
				StatusId:    task.StatusId,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   task.UpdatedAt,
			},
		)
	}

	type fields struct {
		service interfaces.TaskService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nomal",
			fields: fields{
				service: mockService,
			},
			args: args{
				c: c,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl.GetAllTasksController(tt.args.c)
			if rec.Code != http.StatusOK {
				t.Errorf("rec.Code = %v, want %v", rec.Code, http.StatusOK)
			}

			var receivedListResponse response.ListTaskResponse
			err := json.Unmarshal(rec.Body.Bytes(), &receivedListResponse)
			if err != nil {
				t.Errorf("json.Unmarshal() error = %v", err)
				return
			}

			if !cmp.Equal(receivedListResponse.Tasks, expectedListResponse.Tasks) {
				t.Errorf("got %v, want %v", receivedListResponse.Tasks, expectedListResponse.Tasks)
			}
		})
	}
}
