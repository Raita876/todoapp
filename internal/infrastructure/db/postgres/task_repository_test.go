package postgre

import (
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	now = time.Now()
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
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
