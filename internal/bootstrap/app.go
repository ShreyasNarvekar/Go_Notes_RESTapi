package bootstrap

import (
	"go-notes-service/internal/notes"
	"go-notes-service/internal/routes"
	"go-notes-service/internal/tasks"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	frontendDirPath = "./frontend"
)

func NewApp(database *gorm.DB) *fiber.App {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	registerRoutes(v1, database)
	return app
}

func registerRoutes(app fiber.Router, database *gorm.DB) {
	registerHealthAndStatic(app)
	registerNoteRoutes(app, database)
	registerTaskRoutes(app, database)
}

func registerHealthAndStatic(app fiber.Router) {
	app.Get("/health", func(c *fiber.Ctx) error {
		logrus.Info("Health check endpoint hit")
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Static("/", frontendDirPath)
}

func registerNoteRoutes(app fiber.Router, database *gorm.DB) {
	noteRepository := notes.NewNoteRepository(database)
	noteService := notes.NewNoteService(noteRepository)
	noteHandler := notes.NewNoteHandler(noteService)
	routes.NotesRoutes(app, noteHandler)
}

func registerTaskRoutes(app fiber.Router, database *gorm.DB) {
	taskRepository := tasks.NewTaskRepository(database)
	taskService := tasks.NewTaskService(taskRepository)
	taskHandler := tasks.NewTaskHandler(taskService)
	routes.TasksRoutes(app, taskHandler)
}
