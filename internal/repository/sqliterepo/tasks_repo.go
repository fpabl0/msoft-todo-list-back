package sqliterepo

import (
	"errors"

	"github.com/msoft-g1/todo-list-backend/internal/domain/task"
	"gorm.io/gorm"
)

var _ task.Repository = (*tasksRepo)(nil)

// tasksRepo sqlite users repository.
type tasksRepo struct {
	db *gorm.DB
}

// NewTasksRepo creates a new tasks repository.
func NewTasksRepo(db *gorm.DB) task.Repository {
	db.AutoMigrate(&task.Task{})
	return &tasksRepo{db}
}

func (r *tasksRepo) Create(t *task.Task) (*task.Task, error) {
	err := r.db.Create(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *tasksRepo) FindByID(id uint) (*task.Task, error) {
	var t task.Task
	err := r.db.First(&t, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *tasksRepo) FindAll() ([]*task.Task, error) {
	var tasks []*task.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *tasksRepo) FindAllByUser(userID uint) ([]*task.Task, error) {
	var tasks []*task.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *tasksRepo) Update(t *task.Task) (*task.Task, error) {
	err := r.db.Model(&t).Updates(t).Error
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *tasksRepo) Delete(id uint) error {
	return r.db.Delete(&task.Task{}, id).Error
}
