package entities

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	now = time.Now()
)

func TestTask_Validate(t *testing.T) {
	type fields struct {
		Id          uuid.UUID
		Name        string
		Description string
		StatusId    int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
		{
			name: "abnormal",
			fields: fields{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				StatusId:    tt.fields.StatusId,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Task.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTask(t *testing.T) {
	type args struct {
		name        string
		description string
		statusId    int
	}
	tests := []struct {
		name string
		args args
		want *Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTask(tt.args.name, tt.args.description, tt.args.statusId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_UpdateName(t *testing.T) {
	type fields struct {
		Id          uuid.UUID
		Name        string
		Description string
		StatusId    int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	type args struct {
		name string
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
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			args: args{
				name: "Updated task",
			},
			wantErr: false,
		},
		{
			name: "abnormal",
			fields: fields{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			args: args{
				name: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				StatusId:    tt.fields.StatusId,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := tr.UpdateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Task.UpdateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTask_UpdateDescription(t *testing.T) {
	type fields struct {
		Id          uuid.UUID
		Name        string
		Description string
		StatusId    int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	type args struct {
		description string
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
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			args: args{
				description: "Updated description",
			},
			wantErr: false,
		},
		{
			name: "abnormal",
			fields: fields{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			args: args{
				description: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				StatusId:    tt.fields.StatusId,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := tr.UpdateDescription(tt.args.description); (err != nil) != tt.wantErr {
				t.Errorf("Task.UpdateDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTask_UpdateStatus(t *testing.T) {
	type fields struct {
		Id          uuid.UUID
		Name        string
		Description string
		StatusId    int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	type args struct {
		statusId int
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
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			args: args{
				statusId: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				StatusId:    tt.fields.StatusId,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := tr.UpdateStatus(tt.args.statusId); (err != nil) != tt.wantErr {
				t.Errorf("Task.UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEqualTask(t *testing.T) {
	type args struct {
		src *Task
		dst *Task
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal_true",
			args: args{
				src: &Task{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				dst: &Task{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want: true,
		},
		{
			name: "normal_false",
			args: args{
				src: &Task{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
				dst: &Task{
					Id:          uuid.MustParse("fad796a1-e0ed-4ee5-9f88-9b7258d35ae9"),
					Name:        "Task Two",
					Description: "This is the second task",
					StatusId:    2,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualTask(tt.args.src, tt.args.dst); got != tt.want {
				t.Errorf("EqualTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
