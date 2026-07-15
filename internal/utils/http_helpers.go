package utils

import "github.com/gofiber/fiber/v2"

func ParseIDParam(c *fiber.Ctx) (int, error) {
	return c.ParamsInt("id")
}

func WriteError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func WriteErrorMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
