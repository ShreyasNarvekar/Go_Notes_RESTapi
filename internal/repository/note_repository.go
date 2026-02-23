package repository

import "go-notes-service/internal/models"

type NoteRepository interface {
	Create(note *models.Note) error
	GetAll() ([]models.Note, error)
	GetByID(id int) (*models.Note, error)
	Update(note *models.Note) error
	Delete(id int) error
}

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id int) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id int) error
}
