package handlers

import "github.com/gofiber/fiber/v2"

func parseIDParam(c *fiber.Ctx) (int, error) {
	return c.ParamsInt("id")
}

func writeError(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func writeErrorMessage(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"error": message,
	})
}
