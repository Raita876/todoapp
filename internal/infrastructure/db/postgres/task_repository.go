package postgre

import (
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

	if err := repo.db.Find(&dbTasks).Error; err != nil {
		return nil, err
	}

	tasks := make([]*entities.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = fromDBTask(&dbTask)
	}

	return tasks, nil
}

func (repo *GormTaskRepository) Create(task *entities.Task) (*entities.Task, error) {
	dbTask := toDBTask(task)

	if err := repo.db.Create(dbTask).Error; err != nil {
		return nil, err
	}

	return task, nil
}
