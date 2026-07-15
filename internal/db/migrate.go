package db

import (
	"go-notes-service/internal/notes"
	"go-notes-service/internal/tasks"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(database *gorm.DB) {
	if err := database.AutoMigrate(&notes.Note{}, &tasks.Task{}); err != nil {
		log.Fatalf("could not run database migrations: %v", err)
	}
}
