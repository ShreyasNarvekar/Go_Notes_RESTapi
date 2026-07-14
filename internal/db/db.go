package db

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	defaultDBType = "postgres"
	defaultDBHost = "localhost"
	defaultDBName = "notesdb"
)

func envOrDefault(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func buildPostgresDSN() string {
	host := envOrDefault("DB_HOST", defaultDBHost)
	port := envOrDefault("DB_PORT", "5432")
	user := envOrDefault("DB_USER", "postgres")
	password := envOrDefault("DB_PASSWORD", "admin")
	name := envOrDefault("DB_NAME", defaultDBName)

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
}

func buildMySQLDSN() string {
	host := envOrDefault("DB_HOST", defaultDBHost)
	port := envOrDefault("DB_PORT", "3306")
	user := envOrDefault("DB_USER", "user")
	password := envOrDefault("DB_PASSWORD", "password")
	name := envOrDefault("DB_NAME", defaultDBName)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)
}

func Connect() (*gorm.DB, error) {
	dbType := strings.ToLower(strings.TrimSpace(os.Getenv("DB_TYPE")))
	if dbType == "" {
		dbType = defaultDBType
	}

	var dialector gorm.Dialector

	switch dbType {
	case "postgres":
		dialector = postgres.Open(buildPostgresDSN())
	case "mysql":
		dialector = mysql.Open(buildMySQLDSN())
	case "sqlite":
		dialector = sqlite.Open(":memory:")

	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Printf("Database connected using %s\n", dbType)
	return database, nil
}

func Close(database *gorm.DB) {
	sqlDB, err := database.DB()
	if err != nil {
		fmt.Printf("Error getting database connection while closing the database connection: %v\n", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		fmt.Printf("Error closing database connection: %v\n", err)
	} else {
		fmt.Println("Database connection closed")
	}
}
