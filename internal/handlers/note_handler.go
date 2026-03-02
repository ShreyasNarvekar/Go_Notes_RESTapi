package handlers

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type noteHandler struct {
	service services.NoteService
}

func NewNoteHandler(service services.NoteService) NoteHandler {
	return &noteHandler{service: service}
}

// CREATE
func (h *noteHandler) Create(c *fiber.Ctx) error {
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	created, err := h.service.Create(note)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(created)
}

// GET ALL
func (h *noteHandler) GetAll(c *fiber.Ctx) error {
	notes, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(notes)
}

// GET BY ID
func (h *noteHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	note, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(note)
}

// UPDATE
func (h *noteHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updated, err := h.service.Update(id, note)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updated)
}

// DELETE
func (h *noteHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "note deleted successfully",
	})
}
