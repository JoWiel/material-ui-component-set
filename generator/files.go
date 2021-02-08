package generator

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

// CreateDiretoryIfNotExist makes directory if non exists
func CreateDiretoryIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
}

//CreatePathIfNotExsts creates the whole path if it does not exists

// SaveFiles saves all the files to the defined srcDirectory
func SaveFiles(c *fiber.Ctx, files []*multipart.FileHeader, pathPrefix string) error {
	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

		// create directory
		CreateDiretoryIfNotExist(pathPrefix)

		path := pathPrefix + "/" + file.Filename

		CreateDiretoryIfNotExist(path)
		// Save the files to disk:
		err := c.SaveFile(file, fmt.Sprintf("./%s", path))

		// Check for errors
		if err != nil {
			return err
		}
	}
	return nil
}
