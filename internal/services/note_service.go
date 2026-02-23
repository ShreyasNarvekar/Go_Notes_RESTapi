package services

import "go-notes-service/internal/models"

type NoteService interface {
	Create(note models.Note) (models.Note, error)
	GetAll() ([]models.Note, error)
	GetByID(id int) (*models.Note, error)
	Update(id int, note models.Note) (*models.Note, error)
	Delete(id int) error
}
type TaskService interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id int) (*models.Task, error)
	Update(id int, task models.Task) (*models.Task, error)
	Delete(id int) error
}
