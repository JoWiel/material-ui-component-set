package router

import (
	"github.com/JoWiel/component-set-generator/generator"

	"github.com/beevik/guid"
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
		// Parse the multipart form:
		form, err := c.MultipartForm()
		// => *multipart.Form

		// Get all files from "documents" key:
		components := form.File["components"]
		interactions := form.File["interactions"]
		prefabs := form.File["prefabs"]

		// => []*multipart.FileHeader
		newGUID := guid.New()
		directory := "public/uploaded/" + newGUID.String()

		generator.SaveToComponentStore(directory)

		srcDirectory := directory + "/src"

		generator.SaveToComponentStore(srcDirectory)

		err = generator.SaveFiles(c, interactions, srcDirectory + "/interactions")
		err = generator.SaveFiles(c, components, srcDirectory + "/components")
		err = generator.SaveFiles(c, prefabs, srcDirectory + "/prefabs")
		
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		
		go generator.SetGenerator(directory)
		c.Status(200).JSON(&fiber.Map{
			"succes":  true,
			"url": "api/v1/sets/uploaded/" + newGUID.String() + "/dist",
		})
		
		return nil
	})

}