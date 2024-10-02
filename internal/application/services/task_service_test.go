package services

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/application/command"
	"github.com/raita876/todoapp/internal/application/common"
	"github.com/raita876/todoapp/internal/application/interfaces"
	"github.com/raita876/todoapp/internal/application/query"
	"github.com/raita876/todoapp/internal/domain/entities"
	"github.com/raita876/todoapp/internal/domain/repositories"
)

var (
	now = time.Now()
)

type MockTaskRepository struct {
	tasks []*entities.Task
}

func (m *MockTaskRepository) FindAll() ([]*entities.Task, error) {
	return m.tasks, nil
}

func (m *MockTaskRepository) Create(task *entities.Task) (*entities.Task, error) {
	return task, nil
}

func TestNewTaskService(t *testing.T) {
	type args struct {
		taskRepository repositories.TaskRepository
	}
	tests := []struct {
		name string
		args args
		want interfaces.TaskService
	}{
		{
			name: "nomal",
			args: args{
				taskRepository: &MockTaskRepository{
					tasks: []*entities.Task{
						{
							Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
							Name:        "Task One",
							Description: "This is the first task",
							StatusId:    1,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
					},
				},
			},
			want: &TaskService{
				taskRepository: &MockTaskRepository{
					tasks: []*entities.Task{
						{
							Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
							Name:        "Task One",
							Description: "This is the first task",
							StatusId:    1,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskService(tt.args.taskRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_FindAllTasks(t *testing.T) {
	type fields struct {
		taskRepository repositories.TaskRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    *query.TaskQueryListResult
		wantErr bool
	}{
		{
			name: "nomal",
			fields: fields{
				taskRepository: &MockTaskRepository{
					tasks: []*entities.Task{
						{
							Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
							Name:        "Task One",
							Description: "This is the first task",
							StatusId:    1,
							CreatedAt:   now,
							UpdatedAt:   now,
						},
					},
				},
			},
			want: &query.TaskQueryListResult{
				Result: []*common.TaskResult{
					{
						Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
						Name:        "Task One",
						Description: "This is the first task",
						StatusId:    1,
						CreatedAt:   now,
						UpdatedAt:   now,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TaskService{
				taskRepository: tt.fields.taskRepository,
			}
			got, err := ts.FindAllTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.FindAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.FindAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalCreateTaskCommandResult(got, want *command.CreateTaskCommandResult) bool {
	if got == nil && want == nil {
		return true
	}

	if got == nil || want == nil {
		return false
	}

	if got.Result.Name != want.Result.Name ||
		got.Result.Description != want.Result.Description ||
		got.Result.StatusId != want.Result.StatusId {
		return false
	}

	return true
}

func TestTaskService_CreateTask(t *testing.T) {
	type fields struct {
		taskRepository repositories.TaskRepository
	}
	type args struct {
		taskCommand *command.CreateTaskCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.CreateTaskCommandResult
		wantErr bool
	}{
		{
			name: "nomal",
			fields: fields{
				taskRepository: &MockTaskRepository{},
			},
			args: args{
				taskCommand: &command.CreateTaskCommand{
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
				},
			},
			want: &command.CreateTaskCommandResult{
				Result: &common.TaskResult{
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "abnomal empty name",
			fields: fields{
				taskRepository: &MockTaskRepository{},
			},
			args: args{
				taskCommand: &command.CreateTaskCommand{
					Name:        "",
					Description: "This is description",
					StatusId:    1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "abnomal empty description",
			fields: fields{
				taskRepository: &MockTaskRepository{},
			},
			args: args{
				taskCommand: &command.CreateTaskCommand{
					Name:        "This is Name",
					Description: "",
					StatusId:    1,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TaskService{
				taskRepository: tt.fields.taskRepository,
			}
			got, err := ts.CreateTask(tt.args.taskCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !equalCreateTaskCommandResult(got, tt.want) {
				t.Errorf("TaskService.CreateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
