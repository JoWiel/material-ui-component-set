package merger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//PackageJSON is a struct of package.json used for Betty Blocks components-sets same as React
type PackageJSON struct {
	Name            string                 `json:"name"`
	Version         string                 `json:"version"`
	Main            string                 `json:"main"`
	License         string                 `json:"license"`
	Private         bool                   `json:"private"`
	DevDependencies map[string]interface{} `json:"devDependencies"`
	Husky           map[string]interface{} `json:"husky"`
	Scripts         map[string]interface{} `json:"scripts"`
	Files           []string               `json:"files"`
	Dependencies    map[string]interface{} `json:"dependencies"`
	Repository      map[string]interface{} `json:"repository"`
}

//PackageJSONs is a struct of a collection of package.json's used for Betty Blocks components-sets same as React
type PackageJSONs struct {
	PackageJSON []PackageJSON
}

// GetPackageJSONFiles gets the JSONFiles and appends them to one
func GetPackageJSONFiles(files []string) []PackageJSON {
	// result := make(map[string]interface{}, len(files))
	var result []PackageJSON
	for i := range files {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println(err)
		}
		if file == nil {
			fmt.Println(`No such file found: ` + files[i])
		}
		defer file.Close()
		jsonBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		var packageJSON PackageJSON
		// var JSONResult map[string]interface{}
		err = json.Unmarshal([]byte(jsonBytes), &packageJSON)
		result = append(result, packageJSON)
	}
	return result
}

// MergePackages into one the first reporistory given is the leading repository. All unique dependecies will be merged in to one
func MergePackages(newFilePath string, paths []string) error {
	JSONFiles := GetPackageJSONFiles(paths)
	mergedResult := MergePackageJSONs(JSONFiles)

	outputJSON, err := json.Marshal(mergedResult)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(newFilePath, outputJSON, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

// MergePackageJSONs merges the package.json's in to one package.json. All the dependecies are merged if unique in to the first file given
func MergePackageJSONs(packageJSONs []PackageJSON) PackageJSON {
	merged := packageJSONs[0]
	if len(packageJSONs) == 0 {
		fmt.Println(`At least one package must be provided`)
	}
	for i := range packageJSONs {
		dependecies := MergeMaps(merged.Dependencies, packageJSONs[i].Dependencies)
		merged.Dependencies = dependecies
	}
	return merged
}
