package main

import (
	"go-notes-service/internal/handlers"
	"go-notes-service/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	noteService := services.NewNoteService() //Memory based note service
	noteHandler := handlers.NewNoteHandler(noteService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Post("/notes", noteHandler.Create)
	app.Get("/notes", noteHandler.GetAll)
	app.Get("/notes/:id", noteHandler.GetByID)
	app.Put("/notes/:id", noteHandler.Update)
	app.Delete("/notes/:id", noteHandler.Delete)

	log.Fatal(app.Listen(":8080"))
}
