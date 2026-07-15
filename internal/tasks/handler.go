package tasks

import (
	"go-notes-service/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler interface {
	Create(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type taskHandler struct {
	service TaskService
}

func NewTaskHandler(service TaskService) TaskHandler {
	return &taskHandler{
		service: service,
	}
}

func (th *taskHandler) Create(c *fiber.Ctx) error {
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}
	created, err := th.service.Create(task)
	if err != nil {
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}
	return c.Status(fiber.StatusCreated).JSON(created)
}

func (th *taskHandler) GetAll(c *fiber.Ctx) error {
	allTasks, err := th.service.GetAll()
	if err != nil {
		return utils.WriteError(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusOK).JSON(allTasks)
}

func (th *taskHandler) GetByID(c *fiber.Ctx) error {
	id, err := utils.ParseIDParam(c)
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	task, err := th.service.GetByID(id)
	if err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(task)
}

func (th *taskHandler) Update(c *fiber.Ctx) error {
	id, err := utils.ParseIDParam(c)
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	var task Task
	if err = c.BodyParser(&task); err != nil {
		return utils.WriteError(c, fiber.StatusBadRequest, err)
	}
	updated, err := th.service.Update(id, task)
	if err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(updated)
}

func (th *taskHandler) Delete(c *fiber.Ctx) error {
	id, err := utils.ParseIDParam(c)
	if err != nil {
		return utils.WriteErrorMessage(c, fiber.StatusBadRequest, "invalid id")
	}
	if err := th.service.Delete(id); err != nil {
		return utils.WriteError(c, fiber.StatusNotFound, err)
	}
	return c.JSON(fiber.Map{
		"message": "task deleted successfully",
	})
}
