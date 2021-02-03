package middleware

import (
	"github.com/JoWiel/component-set-generator/generator"
	"github.com/beevik/guid"
	"github.com/gofiber/fiber/v2"
)

// UploadSets handles the upload of the component sets
func UploadSets(c *fiber.Ctx) error {
		// Parse the multipart form:
		form, err := c.MultipartForm()
		// => *multipart.Form

		// Get all files from "documents" key:
		components := form.File["components"]
		interactions := form.File["interactions"]
		prefabs := form.File["prefabs"]
		
		if len(components) == 0 {
			c.Status(422).JSON(&fiber.Map{
				"succes":  false,
				"message": "Component(s) must be provided",
			})
		}
		
		if len(prefabs) == 0 {
			c.Status(422).JSON(&fiber.Map{
				"succes":  false,
				"message": "Prefab(s) must be provided",
			})
		}

		// => []*multipart.FileHeader
		newGUID := guid.New()
		directory := "public/uploaded/" + newGUID.String()

		generator.CreateDiretoryIfNotExistint(directory)

		srcDirectory := directory + "/src"

		generator.CreateDiretoryIfNotExistint(srcDirectory)

		err = generator.SaveFiles(c, interactions, srcDirectory+"/interactions")
		err = generator.SaveFiles(c, components, srcDirectory+"/components")
		err = generator.SaveFiles(c, prefabs, srcDirectory+"/prefabs")

		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"succes":  false,
				"message": err,
			})
		}

		go generator.SetGenerator(directory)
		c.Status(200).JSON(&fiber.Map{
			"succes": true,
			"url":    "api/v1/sets/uploaded/" + newGUID.String() + "/dist",
		})

		return nil
	}