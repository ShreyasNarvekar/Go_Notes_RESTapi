package main

import (
	"go-notes-service/internal/db"
	"go-notes-service/internal/handlers"
	"go-notes-service/internal/models"
	"go-notes-service/internal/repository"
	"go-notes-service/internal/routes"
	"go-notes-service/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const (
	serverAddr      = ":8080"
	frontendDirPath = "./frontend"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	mustConnectDB()
	defer db.Close()
	runMigrations()
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	registerRoutes(v1)
	if err := app.Listen(serverAddr); err != nil {
		log.Printf("server stopped with error: %v", err)
	}
}

func mustConnectDB() {
	if err := db.Connect(); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
}

func runMigrations() {
	if err := db.DB.AutoMigrate(&models.Note{}, &models.Task{}); err != nil {
		log.Fatalf("could not run database migrations: %v", err)
	}
}

func registerRoutes(app fiber.Router) {
	registerHealthAndStatic(app)
	registerNoteRoutes(app)
	registerTaskRoutes(app)
}

func registerHealthAndStatic(app fiber.Router) {
	app.Get("/health", func(c *fiber.Ctx) error {
		logrus.Info("Health check endpoint hit")
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Static("/", frontendDirPath)
}

func registerNoteRoutes(app fiber.Router) {
	noteRepository := repository.NewNoteRepository(db.DB)
	noteService := services.NewNoteService(noteRepository)
	noteHandler := handlers.NewNoteHandler(noteService)
	routes.NotesRoutes(app, noteHandler)
}

func registerTaskRoutes(app fiber.Router) {
	taskRepository := repository.NewTaskRepository(db.DB)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	routes.TasksRoutes(app, taskHandler)
}
