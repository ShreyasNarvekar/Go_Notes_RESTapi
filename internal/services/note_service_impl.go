package services

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/repository"
)

type noteService struct {
	repo repository.NoteRepository
}

// NewNoteService creates a new instance of NoteService interface
func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{
		repo: repo,
	}
}

// Create adds a new note to the repository and returns the created note with its ID
func (s *noteService) Create(note models.Note) (models.Note, error) {

	if err := s.repo.Create(&note); err != nil {
		return models.Note{}, err
	}
	return note, nil
}

// GetAll retrieves all notes from the repository and returns them as a slice
func (s *noteService) GetAll() ([]models.Note, error) {

	return s.repo.GetAll()
}

// GetByID retrieves a note by its ID from the repository and returns it
func (s *noteService) GetByID(id int) (*models.Note, error) {
	return s.repo.GetByID(id)
}

// Update modifies an existing note identified by its ID with the provided updated note data and returns the updated note
func (s *noteService) Update(id int, updated models.Note) (*models.Note, error) {
	note, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	note.Title = updated.Title
	note.Content = updated.Content
	if err := s.repo.Update(note); err != nil {
		return nil, err
	}
	return note, nil
}

// Delete removes a note identified by its ID from the repository and returns an error if the operation fails
func (s *noteService) Delete(id int) error {
	return s.repo.Delete(id)
}
