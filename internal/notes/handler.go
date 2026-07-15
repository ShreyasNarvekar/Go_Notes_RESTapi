package notes

import (
	"encoding/json"
	"go-notes-service/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NoteHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type noteHandler struct {
	service NoteService
}

func NewNoteHandler(service NoteService) NoteHandler {
	return &noteHandler{service: service}
}

// CREATE
func (h *noteHandler) Create(c *fiber.Ctx) error {
	logrus.Info("Create note endpoint hit")
	var note Note

	if err := c.BodyParser(&note); err != nil {
		logrus.Errorf("Error parsing request body: %v", err)
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}

	created, err := h.service.Create(note)
	if err != nil {
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (h *noteHandler) GetAll(c *fiber.Ctx) error {
	logrus.Info("Get all notes endpoint hit")
	notes, err := h.service.GetAll()
	if err != nil {
		return utils.WriteError(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusOK).JSON(notes)
}

// GET BY ID
func (h *noteHandler) GetByID(c *fiber.Ctx) error {
	logrus.Info("Get note by ID endpoint hit")
	id, err := c.ParamsInt("id")
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	note, err := h.service.GetByID(id)
	if err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}
	jsonStr := `{"name":"Shreyas","age":26}`
	json.NewEncoder(c).Encode(jsonStr)

	return c.JSON(note)
}

// UPDATE
func (h *noteHandler) Update(c *fiber.Ctx) error {
	logrus.Info("Update note endpoint hit")
	id, err := utils.ParseIDParam(c)
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	var note Note
	if err := c.BodyParser(&note); err != nil {
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}

	updated, err := h.service.Update(id, note)
	if err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(updated)
}

// DELETE
func (h *noteHandler) Delete(c *fiber.Ctx) error {
	logrus.Info("Delete note endpoint hit")
	id, err := utils.ParseIDParam(c)
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.Delete(id); err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}

	return c.JSON(fiber.Map{
		"message": "note deleted successfully",
	})
}
