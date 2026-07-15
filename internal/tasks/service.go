package tasks

import (
	"time"
)

type TaskService interface {
	Create(task Task) (Task, error)
	GetAll() ([]Task, error)
	GetByID(id int) (*Task, error)
	Update(id int, task Task) (*Task, error)
	Delete(id int) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (ts *taskService) Create(task Task) (Task, error) {
	if err := ts.repo.Create(&task); err != nil {
		return Task{}, err
	}
	return task, nil
}
func (ts *taskService) GetAll() ([]Task, error) {
	return ts.repo.GetAll()
}
func (ts *taskService) GetByID(id int) (*Task, error) {
	return ts.repo.GetByID(id)
}
func (ts *taskService) Update(id int, updated Task) (*Task, error) {
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
