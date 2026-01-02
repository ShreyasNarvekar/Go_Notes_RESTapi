package services

import (
	"errors"
	"go-notes-service/internal/models"
)

type noteService struct {
	notes []models.Note
}

func NewNoteService() NoteService { //Here we are returning the interface NoteService
	return &noteService{
		notes: []models.Note{},
	}
}

func (s *noteService) Create(note models.Note) models.Note {
	note.ID = len(s.notes) + 1
	s.notes = append(s.notes, note)
	return note // is this correct?
}

func (s *noteService) GetAll() []models.Note {
	return s.notes
}

func (s *noteService) GetByID(id int) (*models.Note, error) {
	for _, note := range s.notes {
		if note.ID == id {
			return &note, nil
		}
	}
	return nil, errors.New("note not found")
}

func (s *noteService) Update(id int, updated models.Note) (*models.Note, error) {
	for i, note := range s.notes {
		if note.ID == id {
			s.notes[i].Title = updated.Title
			s.notes[i].Content = updated.Content
			return &s.notes[i], nil
		}
	}
	return nil, errors.New("note not found")
}

func (s *noteService) Delete(id int) error {
	for i, note := range s.notes {
		if note.ID == id {
			s.notes = append(s.notes[:i], s.notes[i+1:]...)
			return nil
		}
	}
	return errors.New("note not found")
}
