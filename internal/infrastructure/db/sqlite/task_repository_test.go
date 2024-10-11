package sqlite

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/raita876/todoapp/internal/domain/entities"
	"github.com/raita876/todoapp/internal/infrastructure/db/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	now = time.Now()
)

func setupDatabase() (*gorm.DB, func(), error) {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, nil, errors.New("Failed to connect to database")
	}

	err = database.AutoMigrate(&postgres.Task{})
	if err != nil {
		return nil, nil, errors.New("Failed to migrate database")
	}

	cleanup := func() {
		database.Exec("DELETE FROM tasks")
	}

	return database, cleanup, nil
}

func TestGormTaskRepository_FindAll(t *testing.T) {
	database, cleanup, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := postgres.NewGormTaskRepository(database)

	_, err = repo.Create(&entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Task One",
		Description: "This is the first task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Create(&entities.Task{
		Id:          uuid.MustParse("fad796a1-e0ed-4ee5-9f88-9b7258d35ae9"),
		Name:        "Task Two",
		Description: "This is the second task",
		StatusId:    2,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := repo.FindAll()
	if err != nil || len(tasks) != 2 {
		t.Error("Error fetching all tasks or count mismatch")
	}
}

func TestGormTaskRepository_Create(t *testing.T) {
	database, cleanup, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := postgres.NewGormTaskRepository(database)

	_, err = repo.Create(&entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Task One",
		Description: "This is the first task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Errorf("failed create task: %s", err)
	}
}

func TestGormTaskRepository_Update(t *testing.T) {
	database, cleanup, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := postgres.NewGormTaskRepository(database)

	_, err = repo.Create(&entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Task One",
		Description: "This is the first task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Errorf("failed create task: %s", err)
	}

	want := &entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Updated Task",
		Description: "This is updated task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	got, err := repo.Update(want)
	if err != nil {
		t.Errorf("failed update task: %s", err)
	}

	// update によりタイムスタンプが更新されるため、 time.Time の比較は除外
	if !cmp.Equal(got.Id, want.Id) {
		t.Errorf("got %v, want %v", got, want)
	}
	if !cmp.Equal(got.Name, want.Name) {
		t.Errorf("got %v, want %v", got, want)
	}
	if !cmp.Equal(got.Description, want.Description) {
		t.Errorf("got %v, want %v", got, want)
	}
	if !cmp.Equal(got.StatusId, want.StatusId) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGormTaskRepository_Delete(t *testing.T) {
	database, cleanup, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := postgres.NewGormTaskRepository(database)

	_, err = repo.Create(&entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Task One",
		Description: "This is the first task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Errorf("failed create task: %s", err)
	}

	err = repo.Delete(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"))
	if err != nil {
		t.Errorf("failed delete task: %s", err)
	}
}

func TestGormTaskRepository_FindTaskById(t *testing.T) {
	database, cleanup, err := setupDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := postgres.NewGormTaskRepository(database)

	want, err := repo.Create(&entities.Task{
		Id:          uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"),
		Name:        "Task One",
		Description: "This is the first task",
		StatusId:    1,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		t.Errorf("failed create task: %s", err)
	}

	got, err := repo.FindTaskById(uuid.MustParse("b81240b0-7122-4d06-bdb2-8bcf512d6c63"))
	if err != nil {
		t.Errorf("failed find task: %s", err)
	}

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
