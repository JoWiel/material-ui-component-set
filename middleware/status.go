package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// SendErrorMessage sends the error message in JSON format
func SendErrorMessage(c *fiber.Ctx, statusCode int, err interface{}) {
	c.Status(statusCode).JSON(&fiber.Map{
		"succes": false,
		"message":  err,
	})
}