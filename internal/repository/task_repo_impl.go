package repository

import (
	"go-notes-service/internal/models"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (tr *taskRepository) Create(task *models.Task) error {
	return tr.db.Create(task).Error
}
func (tr *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := tr.db.Find(&tasks).Error
	return tasks, err
}
func (tr *taskRepository) GetByID(id int) (*models.Task, error) {
	var task models.Task
	err := tr.db.First(&task, id).Error
	return &task, err
}
func (tr *taskRepository) Update(task *models.Task) error {
	return tr.db.Save(task).Error
}
func (tr *taskRepository) Delete(id int) error {
	return tr.db.Delete(id).Error
}
