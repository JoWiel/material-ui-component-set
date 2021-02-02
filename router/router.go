package router

import (
	"fmt"
	"os"

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
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}
		// => *multipart.Form

		// Get all files from "documents" key:
		components := form.File["components"]
		interactions := form.File["interactions"]
		prefabs := form.File["prefabs"]

		// => []*multipart.FileHeader
		newGUID := guid.New()
		directory := "public/uploaded/" + newGUID.String()

		if _, err := os.Stat(directory); os.IsNotExist(err) {
			os.Mkdir(directory, 0700)
		}

		srcDirectory := directory + "/src"

		if _, err := os.Stat(srcDirectory); os.IsNotExist(err) {
			os.Mkdir(srcDirectory, 0700)
		}

		// generator.StaticGenerator(directory)
		// Loop through files:
		for _, file := range components {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// create directory
			pathPrefix := srcDirectory + "/components"

			if _, err := os.Stat(pathPrefix); os.IsNotExist(err) {
				os.Mkdir(pathPrefix, 0700)
			}

			path := pathPrefix + "/" + file.Filename

			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", path))

			// Check for errors
			if err != nil {
				c.Status(500).JSON(&fiber.Map{
					"succes":  false,
					"message": err,
				})
			}
		}

		for _, file := range interactions {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// create directory
			pathPrefix := srcDirectory + "/interactions"

			if _, err := os.Stat(pathPrefix); os.IsNotExist(err) {
				os.Mkdir(pathPrefix, 0700)
			}

			path := pathPrefix + "/" + file.Filename
			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", path))

			// Check for errors
			if err != nil {
				c.Status(500).JSON(&fiber.Map{
					"succes":  false,
					"message": err,
				})
			}
		}

		for _, file := range prefabs {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// create directory
			pathPrefix := srcDirectory + "/prefabs"

			if _, err := os.Stat(pathPrefix); os.IsNotExist(err) {
				os.Mkdir(pathPrefix, 0700)
			}

			path := pathPrefix + "/" + file.Filename
			// Save the files to disk:
			err := c.SaveFile(file, fmt.Sprintf("./%s", path))

			// Check for errors
			if err != nil {
				c.Status(500).JSON(&fiber.Map{
					"succes":  false,
					"message": err,
				})
			}
		}
		go generator.SetGenerator(directory)
		
		c.Status(200).JSON((&fiber.Map{
			"succes":  true,
			"url": "/api/vq/sets/uploaded/" + newGUID.String(),
		}))
		
		return nil
	})

}

func saveToComponentStore() {

}
