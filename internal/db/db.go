package db

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	defaultDBType = "postgres"
	postgresDSN   = "host=localhost user=postgres password=admin dbname=notesdb port=5432 sslmode=disable"
	mysqlDSN      = "user:password@tcp(localhost:3306)/notesdb?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() error {
	dbType := strings.ToLower(strings.TrimSpace(os.Getenv("DB_TYPE")))
	if dbType == "" {
		dbType = defaultDBType
	}

	var dialector gorm.Dialector

	switch dbType {
	case "postgres":
		dialector = postgres.Open(postgresDSN)
	case "mysql":
		dialector = mysql.Open(mysqlDSN)
	default:
		return fmt.Errorf("unsupported database type: %s", dbType)
	}

	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = database
	fmt.Printf("Database connected using %s\n", dbType)
	return nil
}
