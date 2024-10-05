package postgre

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/domain/entities"
	"github.com/raita876/todoapp/internal/domain/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	now = time.Now()
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func TestNewGormTaskRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want repositories.TaskRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGormTaskRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGormTaskRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGormTaskRepository_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		want    []*entities.Task
		wantErr bool
	}{
		{
			name: "nomal",
			want: []*entities.Task{
				{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := NewDbMock()
			if err != nil {
				t.Errorf("Failed to initialize mock DB: %v", err)
			}

			rows := sqlmock.NewRows([]string{"id", "name", "description", "status_id", "created_at", "updated_at"}).
				AddRow(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"), "Task One", "This is the first task", 1, now, now)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks"`)).
				WillReturnRows(rows)

			repo := &GormTaskRepository{
				db: mockDB,
			}
			got, err := repo.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GormTaskRepository.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GormTaskRepository.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGormTaskRepository_Create(t *testing.T) {

	type args struct {
		task *entities.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Task
		wantErr bool
	}{
		{
			name: "nomal",
			args: args{
				task: &entities.Task{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want: &entities.Task{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := NewDbMock()
			if err != nil {
				t.Errorf("Failed to initialize mock DB: %v", err)
			}

			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "tasks"`)).
				WithArgs(tt.args.task.Name, tt.args.task.Description, tt.args.task.StatusId, tt.args.task.CreatedAt, tt.args.task.UpdatedAt, tt.args.task.Id).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow((tt.args.task.Id)))
			mock.ExpectCommit()

			repo := &GormTaskRepository{
				db: mockDB,
			}
			got, err := repo.Create(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("GormTaskRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GormTaskRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: 別パッケージに移植
func equalTask(got, want *entities.Task) bool {
	if got == nil && want == nil {
		return true
	}

	if got == nil || want == nil {
		return false
	}

	if got.Name != want.Name ||
		got.Description != want.Description ||
		got.StatusId != want.StatusId {
		return false
	}

	return true
}

func TestGormTaskRepository_Update(t *testing.T) {
	type args struct {
		task *entities.Task
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Task
		wantErr bool
	}{
		{
			name: "nomal",
			args: args{
				task: &entities.Task{
					Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
					Name:        "Task One",
					Description: "This is the first task",
					StatusId:    1,
					CreatedAt:   now,
					UpdatedAt:   now,
				},
			},
			want: &entities.Task{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := NewDbMock()
			if err != nil {
				t.Errorf("Failed to initialize mock DB: %v", err)
			}

			mock.ExpectBegin()
			mock.ExpectExec("UPDATE \"tasks\"").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			rows := sqlmock.NewRows([]string{"id", "name", "description", "status_id", "created_at", "updated_at"}).
				AddRow(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"), "Task One", "This is the first task", 1, now, now)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks"`)).
				WithArgs(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"), 1).
				WillReturnRows(rows)

			repo := &GormTaskRepository{
				db: mockDB,
			}

			got, err := repo.Update(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("GormTaskRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalTask(got, tt.want) {
				t.Errorf("GormTaskRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGormTaskRepository_Delete(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nomal",
			args: args{
				id: uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := NewDbMock()
			if err != nil {
				t.Errorf("Failed to initialize mock DB: %v", err)
			}

			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM \"tasks\"").
				WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			repo := &GormTaskRepository{
				db: mockDB,
			}
			if err := repo.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GormTaskRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGormTaskRepository_FindById(t *testing.T) {
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Task
		wantErr bool
	}{
		{
			name: "nomal",
			args: args{
				id: uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
			},
			want: &entities.Task{
				Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
				Name:        "Task One",
				Description: "This is the first task",
				StatusId:    1,
				CreatedAt:   now,
				UpdatedAt:   now,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := NewDbMock()
			if err != nil {
				t.Errorf("Failed to initialize mock DB: %v", err)
			}

			rows := sqlmock.NewRows([]string{"id", "name", "description", "status_id", "created_at", "updated_at"}).
				AddRow(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"), "Task One", "This is the first task", 1, now, now)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks"`)).
				WithArgs(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"), 1).
				WillReturnRows(rows)

			repo := &GormTaskRepository{
				db: mockDB,
			}
			got, err := repo.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GormTaskRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GormTaskRepository.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}
