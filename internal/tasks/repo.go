package tasks

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *Task) error
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task *Task) error {
	return r.db.Create(task).Error
}
func (r *taskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) GetByID(id int) (*Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	return &task, err
}
func (r *taskRepository) Update(task *Task) error {
	return r.db.Save(task).Error
}
func (r *taskRepository) Delete(id int) error {
	return r.db.Delete(&Task{}, id).Error
}
