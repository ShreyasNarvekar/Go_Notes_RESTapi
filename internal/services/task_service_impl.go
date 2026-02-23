package services

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/repository"
	"time"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (ts *taskService) Create(task models.Task) (models.Task, error) {
	if err := ts.repo.Create(&task); err != nil {
		return models.Task{}, err
	}
	return task, nil
}
func (ts *taskService) GetAll() ([]models.Task, error) {
	return ts.repo.GetAll()
}
func (ts *taskService) GetByID(id int) (*models.Task, error) {
	return ts.repo.GetByID(id)
}
func (ts *taskService) Update(id int, updated models.Task) (*models.Task, error) {
	task, err := ts.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	task.Title = updated.Title
	task.Description = updated.Description
	task.Completed = updated.Completed
	task.UpdatedAt = time.Now()
	if err = ts.repo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}
func (ts *taskService) Delete(id int) error {
	return ts.repo.Delete(id)
}
