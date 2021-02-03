package generator

import (
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

// SetGenerator generates custom component sets from the selected sets
func SetGenerator(path string) {
	projectRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	cmdLine := projectRoot + "/generator/generator.sh"

	// line := "cd " + projectRoot + " && pwd"
	command := exec.Command("/bin/sh", cmdLine, path)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Run the command
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("command succesfully ran:" + cmdLine)
}

// StaticGenerator copies the static files
func StaticGenerator(path string) {
	projectRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	cmdLine := projectRoot + "/generator/static.sh"

	// line := "cd " + projectRoot + " && pwd"
	command := exec.Command("/bin/sh", cmdLine, path)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Run the command
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("command succesfully ran:")
}

// CreateDiretoryIfNotExistint makes directory if non exists
func CreateDiretoryIfNotExistint(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
}

// SaveFiles saves all the files to the defined srcDirectory
func SaveFiles(c *fiber.Ctx, files []*multipart.FileHeader, pathPrefix string) error {
	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

		// create directory
		CreateDiretoryIfNotExistint(pathPrefix)

		path := pathPrefix + "/" + file.Filename

		CreateDiretoryIfNotExistint(path)
		// Save the files to disk:
		err := c.SaveFile(file, fmt.Sprintf("./%s", path))

		// Check for errors
		if err != nil {
			return err
		}
	}
	return nil
}
