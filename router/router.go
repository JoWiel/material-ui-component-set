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
	fileServer := api.Group(`/sets`)
	fileServer.Static("/uploaded/", "./public")

	fileServer.Static("/merged", "./public/build")

	api.Post("/upload", func(c *fiber.Ctx) error {
		// err := middleware.UploadSets(c)
		// if err != nil {
		// 	middleware.SendErrorMessage(c, 500, err)
		// }
		middleware.SendErrorMessage(c, 404, `Disabled`)
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
