package notes

import (
	"gorm.io/gorm"
)

type NoteRepository interface {
	Create(note *Note) error
	GetAll() ([]Note, error)
	GetByID(id int) (*Note, error)
	Update(note *Note) error
	Delete(id int) error
}

type noteRepo struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepo{
		db: db,
	}
}

// Create
func (r *noteRepo) Create(note *Note) error {
	return r.db.Create(note).Error
}

// GetAll
func (r *noteRepo) GetAll() ([]Note, error) {
	var notes []Note
	err := r.db.Find(&notes).Error
	return notes, err
}

// GetByID
func (r *noteRepo) GetByID(id int) (*Note, error) {
	var note Note
	err := r.db.First(&note, id).Error
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// Update
func (r *noteRepo) Update(note *Note) error {
	return r.db.Save(note).Error
}

// Delete
func (r *noteRepo) Delete(id int) error {
	return r.db.Delete(&Note{}, id).Error
}
