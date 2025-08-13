package handler

import "github.com/gofiber/fiber/v2"

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	type badReq interface{ Error() string }
	if e, ok := err.(badReq); ok && e.Error() == "invalid credentials" {
		code = fiber.StatusUnauthorized
	}
	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
