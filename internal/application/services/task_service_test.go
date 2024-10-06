package services

import (
	"errors"
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

func (m *MockTaskRepository) FindTaskById(id uuid.UUID) (*entities.Task, error) {
	for _, t := range m.tasks {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskRepository) Create(task *entities.Task) (*entities.Task, error) {
	m.tasks = append(m.tasks, task)
	return task, nil
}

func (m *MockTaskRepository) Update(task *entities.Task) (*entities.Task, error) {
	for i, t := range m.tasks {
		if t.Id == task.Id {
			m.tasks[i] = task
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (m *MockTaskRepository) Delete(id uuid.UUID) error {
	for i, t := range m.tasks {
		if t.Id == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
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
			name: "normal",
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
			name: "normal",
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
			name: "normal",
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
			name: "abnormal empty name",
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
			name: "abnormal empty description",
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

func TestTaskService_FindTaskById(t *testing.T) {
	type fields struct {
		taskRepository repositories.TaskRepository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *query.TaskQueryResult
		wantErr bool
	}{
		{
			name: "normal",
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
			args: args{
				id: uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
			},
			want: &query.TaskQueryResult{
				Result: &common.TaskResult{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			wantErr: false,
		},
		{
			name: "abnormal",
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
			args: args{
				id: uuid.MustParse("fad796a1-e0ed-4ee5-9f88-9b7258d35ae9"),
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
			got, err := ts.FindTaskById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.FindTaskById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskService.FindTaskById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalUpdateTaskCommandResult(got, want *command.UpdateTaskCommandResult) bool {
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

func TestTaskService_UpdateTask(t *testing.T) {
	type fields struct {
		taskRepository repositories.TaskRepository
	}
	type args struct {
		updateCommand *command.UpdateTaskCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *command.UpdateTaskCommandResult
		wantErr bool
	}{
		{
			name: "normal",
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
			args: args{
				updateCommand: &command.UpdateTaskCommand{
					ID:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Updated task",
					Description: "This is updated task",
					StatusId:    2,
				},
			},
			want: &command.UpdateTaskCommandResult{
				Result: &common.TaskResult{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Updated task",
					Description: "This is updated task",
					StatusId:    2,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			wantErr: false,
		},
		{
			name: "abnormal",
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
			args: args{
				updateCommand: &command.UpdateTaskCommand{
					ID:          uuid.MustParse("fad796a1-e0ed-4ee5-9f88-9b7258d35ae9"),
					Name:        "Updated task",
					Description: "This is updated task",
					StatusId:    2,
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
			got, err := ts.UpdateTask(tt.args.updateCommand)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskService.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalUpdateTaskCommandResult(got, tt.want) {
				t.Errorf("TaskService.UpdateTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	type fields struct {
		taskRepository repositories.TaskRepository
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal",
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
			args: args{
				id: uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
			},
			wantErr: false,
		},
		{
			name: "abnormal",
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
			args: args{
				id: uuid.MustParse("fad796a1-e0ed-4ee5-9f88-9b7258d35ae9"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &TaskService{
				taskRepository: tt.fields.taskRepository,
			}
			if err := ts.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("TaskService.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
