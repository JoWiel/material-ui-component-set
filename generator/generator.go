package generator

import (
	"fmt"
	"os"
	"os/exec"
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

//BuildSet generates custom component sets from the merged sets
func BuildSet(source, destination string) error {
	projectRoot, err := os.Getwd()
	if err != nil {
		return err
	}
	cmdLine := projectRoot + "/generator/build-set.sh"

	sourceCommand := source
	command := exec.Command("/bin/sh", cmdLine, projectRoot, sourceCommand, destination)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	// Run the command
	if err := command.Run(); err != nil {
		return err
	}

	return nil
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
