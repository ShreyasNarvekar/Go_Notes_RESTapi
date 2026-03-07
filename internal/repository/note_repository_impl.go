package repository

import (
	"go-notes-service/internal/models"

	"gorm.io/gorm"
)

type noteRepo struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepo{
		db: db,
	}
}

// Create
func (r *noteRepo) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

// GetAll
func (r *noteRepo) GetAll() ([]models.Note, error) {
	var notes []models.Note
	err := r.db.Find(&notes).Error
	return notes, err
}

// GetByID
func (r *noteRepo) GetByID(id int) (*models.Note, error) {
	var note models.Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// Update
func (r *noteRepo) Update(note *models.Note) error {
	return r.db.Save(note).Error
}

// Delete
func (r *noteRepo) Delete(id int) error {
	return r.db.Delete(&models.Note{}, id).Error
}
