package merger

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/JoWiel/component-set-generator/generator"
)

//Copy copies the file from source to destination
func Copy(source, destination string) error {
	// Read all content of src to data
	data, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//CopyDirectory merges the directp
func CopyDirectory(source, destination string) error {

	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()
		newFileSource := source + `/` + fileName
		newDestination := destination + `/` + fileName
		if strings.Contains(fileName, `.`) {
			err := Copy(newFileSource, newDestination)
			if err != nil {
				return err
			}
		} else {
			newSource := source + `/` + fileName
			generator.CreateDiretoryIfNotExist(newDestination)
			err := CopyDirectory(newSource, newDestination)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//CopyDirectories copies the contents in the directories in to one directory
func CopyDirectories(sources []string, destination string) error {
	for _, source := range sources {
		err := CopyDirectory(source, destination)
		if err != nil {
			return err
		}
	}
	return nil
}
