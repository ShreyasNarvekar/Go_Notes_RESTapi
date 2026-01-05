package services

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/repository"
)

type noteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService { //Here we are returning the interface NoteService
	return &noteService{
		repo: repo,
	}
}

func (s *noteService) Create(note models.Note) (models.Note, error) {
	s.repo.Create(&note)
	return note, nil
}

func (s *noteService) GetAll() ([]models.Note, error) {

	return s.repo.GetAll()
}

func (s *noteService) GetByID(id int) (*models.Note, error) {
	return s.repo.GetByID(id)
}

func (s *noteService) Update(id int, updated models.Note) (*models.Note, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	note.Title = updated.Title
	note.Content = updated.Content
	s.repo.Update(note)
	return note, nil
}

func (s *noteService) Delete(id int) error {
	return s.repo.Delete(id)
}
