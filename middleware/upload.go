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

	if form == nil {
		c.Status(422).JSON(&fiber.Map{
			"succes":  false,
			"message": "Component and prefabs(s) must be provided",
		})
		return nil
	}
	// Get all files from "documents" key:
	components, componentExists := form.File["components"]
	if !componentExists {
		c.Status(422).JSON(&fiber.Map{
			"succes":  false,
			"message": "Component(s) must be provided",
		})
		return nil
	}

	interactions := form.File["interactions"]
	prefabs, prefabsExists := form.File["prefabs"]
	if !prefabsExists {
		c.Status(422).JSON(&fiber.Map{
			"succes":  false,
			"message": "Prefabs(s) must be provided",
		})
		return nil
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
