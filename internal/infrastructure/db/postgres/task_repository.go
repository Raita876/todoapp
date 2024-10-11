package postgres

import (
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

	if err := repo.db.Find(&dbTasks).Error; err != nil {
		return nil, err
	}

	tasks := make([]*entities.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		tasks[i] = fromDBTask(&dbTask)
	}

	return tasks, nil
}

func (repo *GormTaskRepository) FindTaskById(id uuid.UUID) (*entities.Task, error) {
	var dbTask Task
	if err := repo.db.Preload("Status").First(&dbTask, id).Error; err != nil {
		return nil, err
	}

	return fromDBTask(&dbTask), nil
}

func (repo *GormTaskRepository) Create(task *entities.Task) (*entities.Task, error) {
	dbTask := toDBTask(task)

	if err := repo.db.Create(dbTask).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (repo *GormTaskRepository) Update(task *entities.Task) (*entities.Task, error) {
	dbTask := toDBTask(task)
	err := repo.db.Model(&dbTask).Where("id = ?", dbTask.Id).Updates(dbTask).Error
	if err != nil {
		return nil, err
	}

	return repo.FindTaskById(dbTask.Id)
}

func (repo *GormTaskRepository) Delete(id uuid.UUID) error {
	// id で検索して存在しない場合はエラー
	_, err := repo.FindTaskById(id)
	if err != nil {
		return err
	}

	return repo.db.Delete(&Task{}, id).Error
}
