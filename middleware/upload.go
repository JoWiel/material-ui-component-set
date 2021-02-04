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
		SendErrorMessage(c, 422, "Component(s) and prefabs(s) must be provided")
		return nil
	}
	// Get all files from "documents" key:
	components, componentExists := form.File["components"]
	if !componentExists {
		SendErrorMessage(c, 422, "Component(s) must be provided")
		return nil
	}

	interactions := form.File["interactions"]
	prefabs, prefabsExists := form.File["prefabs"]
	if !prefabsExists {
		SendErrorMessage(c, 422, "Prefabs(s) must be provided")
		return nil
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
		SendErrorMessage(c, 500, err)
	}

	go generator.SetGenerator(directory)
	c.Status(200).JSON(&fiber.Map{
		"succes": true,
		"url":    "api/v1/sets/uploaded/" + newGUID.String() + "/dist",
	})

	return nil
}
