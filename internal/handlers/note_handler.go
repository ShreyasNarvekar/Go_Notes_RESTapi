package handlers

import (
	"go-notes-service/internal/models"
	"go-notes-service/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type noteHandler struct {
	service services.NoteService
}

func NewNoteHandler(service services.NoteService) NoteHandler {
	return &noteHandler{service: service}
}

// CREATE
func (h *noteHandler) Create(c *fiber.Ctx) error {
	logrus.Info("Create note endpoint hit")
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		logrus.Errorf("Error parsing request body: %v", err)
		return writeError(c, fiber.StatusBadRequest, err)
	}

	created, err := h.service.Create(note)
	if err != nil {
		return writeError(c, fiber.StatusBadRequest, err)
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h *noteHandler) GetAll(c *fiber.Ctx) error {
	logrus.Info("Get all notes endpoint hit")
	notes, err := h.service.GetAll()
	if err != nil {
		return writeError(c, fiber.StatusInternalServerError, err)
	}
	return c.JSON(notes)
}

// GET BY ID
func (h *noteHandler) GetByID(c *fiber.Ctx) error {
	logrus.Info("Get note by ID endpoint hit")
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	note, err := h.service.GetByID(id)
	if err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(note)
}

// UPDATE
func (h *noteHandler) Update(c *fiber.Ctx) error {
	logrus.Info("Update note endpoint hit")
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	var note models.Note
	if err := c.BodyParser(&note); err != nil {
		return writeError(c, fiber.StatusBadRequest, err)
	}

	updated, err := h.service.Update(id, note)
	if err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(updated)
}

// DELETE
func (h *noteHandler) Delete(c *fiber.Ctx) error {
	logrus.Info("Delete note endpoint hit")
	id, err := parseIDParam(c)
	if err != nil {
		return writeErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		return writeError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(fiber.Map{
		"message": "note deleted successfully",
	})
}
