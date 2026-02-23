package main

import (
	"go-notes-service/internal/db"
	"go-notes-service/internal/handlers"
	"go-notes-service/internal/models"
	"go-notes-service/internal/repository"
	"go-notes-service/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//We need to connect to DB here
	if err := db.Connect(); err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	} //connect to the database
	db.DB.AutoMigrate(&models.Note{}) //AutoMigreate creates table in database if its not present.
	noteRepository := repository.NewNoteRepository(db.DB)
	noteService := services.NewNoteService(noteRepository) //Memory based note service
	noteHandler := handlers.NewNoteHandler(noteService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	app.Post("/notes", noteHandler.Create)
	app.Get("/notes", noteHandler.GetAll)
	app.Get("/notes/:id", noteHandler.GetByID)
	app.Put("/notes/:id", noteHandler.Update)
	app.Delete("/notes/:id", noteHandler.Delete)

	taskRepository := repository.NewTaskRepository()
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	app.Get("/tasks", taskHandler.GetAll)
	app.Get("/tasks/:id", taskHandler.GetByID)
	app.Post("/tasks", taskHandler.Create)
	app.Put("/tasks/:id", taskHandler.UpdateTask)
	app.Delete("/tasks/:id", taskHandler.Delete)

	log.Fatal(app.Listen(":8080"))
}
