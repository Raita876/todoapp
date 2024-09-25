package postgre

import (
	"time"

	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/domain/entities"
	"github.com/raita876/todoapp/internal/domain/repositories"
	"gorm.io/gorm"
)

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &GormTaskRepository{db: db}
}

func (repo *GormTaskRepository) FindAll() ([]*entities.Task, error) {
	var dbTasks []Task

	// TODO: DB の値を dbTasks に格納する
	// 仮のデータを挿入
	dbTasks = append(dbTasks, Task{
		Id:          uuid.New(),
		Name:        "sample name",
		Description: "sample description",
		StatusId:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	tasks := make([]*entities.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = fromDBTask(&dbTask)
	}

	return tasks, nil
}
