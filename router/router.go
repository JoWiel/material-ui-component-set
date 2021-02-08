package router

import (
	"github.com/JoWiel/component-set-generator/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api/v1", logger.New())
	//File server for generated sets
	api.Static("/sets", "./public")

	api.Post("/upload", func(c *fiber.Ctx) error {
		err := middleware.UploadSets(c)
		if err != nil {
			middleware.SendErrorMessage(c, 500, err)
		}
		return nil
	})

	api.Post("/new-set", func(c *fiber.Ctx) error {
		err := middleware.GenerateCombinedSet(c)
		if err != nil {
			middleware.SendErrorMessage(c, 500, err)
		}
		return nil
	})
}
