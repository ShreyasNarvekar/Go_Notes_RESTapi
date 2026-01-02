package handlers

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteHandler struct {
	service services.NoteService
}

func NewNoteHandler(service services.NoteService) *NoteHandler {
	return &NoteHandler{service: service}
}

// CREATE
func (h *NoteHandler) Create(c *fiber.Ctx) error {
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	created := h.service.Create(note)
	return c.JSON(created)
}

// GET ALL
func (h *NoteHandler) GetAll(c *fiber.Ctx) error {
	return c.JSON(h.service.GetAll())
}

// GET BY ID
func (h *NoteHandler) GetByID(c *fiber.Ctx) error {
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
func (h *NoteHandler) Update(c *fiber.Ctx) error {
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
func (h *NoteHandler) Delete(c *fiber.Ctx) error {
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
