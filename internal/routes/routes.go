package routes

import (
	"go-notes-service/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// NotesRoutes defines the routes for note-related operations
func NotesRoutes(app *fiber.App, noteHandler handlers.NoteHandler) {
	app.Post("/notes", noteHandler.Create)
	app.Get("/notes", noteHandler.GetAll)
	app.Get("/notes/:id", noteHandler.GetByID)
	app.Put("/notes/:id", noteHandler.Update)
	app.Delete("/notes/:id", noteHandler.Delete)
}

// TasksRoutes defines the routes for task-related operations
func TasksRoutes(app *fiber.App, taskHandler handlers.TaskHandler) {
	app.Get("/tasks", taskHandler.GetAll)
	app.Get("/tasks/:id", taskHandler.GetByID)
	app.Post("/tasks", taskHandler.Create)
	app.Put("/tasks/:id", taskHandler.Update)
	app.Delete("/tasks/:id", taskHandler.Delete)
}
