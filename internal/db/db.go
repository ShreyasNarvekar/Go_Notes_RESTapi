package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=localhost user=postgres password=admin dbname=notesdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("Failed to connect to database: %w", err)
	}

	DB = database
	fmt.Println("Database connected")
	return nil
}
