package handlers

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/services"

	"github.com/gofiber/fiber/v2"
)

type taskHandler struct {
	service services.TaskService
}

func NewTaskHandler(service services.TaskService) TaskHandler {
	return &taskHandler{
		service: service,
	}
}

func (th *taskHandler) Create(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return writeError(c, fiber.StatusBadRequest, err)
	}
	created, err := th.service.Create(task)
	if err != nil {
		return writeError(c, fiber.StatusBadRequest, err)
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (th *taskHandler) GetAll(c *fiber.Ctx) error {
	allTasks, err := th.service.GetAll()
	if err != nil {
		return writeError(c, fiber.StatusInternalServerError, err)
	}
	return c.JSON(allTasks)
}

func (th *taskHandler) GetByID(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	task, err := th.service.GetByID(id)
	if err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(task)
}

func (th *taskHandler) Update(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	var task models.Task
	if err = c.BodyParser(&task); err != nil {
		return writeError(c, fiber.StatusBadRequest, err)
	}
	updated, err := th.service.Update(id, task)
	if err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(updated)
}

func (th *taskHandler) Delete(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	if err := th.service.Delete(id); err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(fiber.Map{
		"message": "task deleted successfully",
	})
}
