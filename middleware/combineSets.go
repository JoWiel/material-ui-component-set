package middleware

import (
	"github.com/JoWiel/component-set-generator/generator"
	"github.com/JoWiel/component-set-generator/merger"
	"github.com/gofiber/fiber/v2"
)

//RequestBody is the response body struct for the new sets
type RequestBody struct {
	Sets         []string `json:"sets"`
	Organisation string   `json:"organisation"`
	App          string   `json:"app"`
}

func getPaths(requestBody *RequestBody) ([]string, []string, []string) {
	pathToRepositories := `./public/repositories/`
	repositoryPathToIndexJS := `src/`
	var webpackPaths []string
	var packagePaths []string
	var indexJSPaths []string

	for i := range requestBody.Sets {
		prefix := pathToRepositories + requestBody.Sets[i] + `/`
		webpackPaths = append(webpackPaths, (prefix + `webpack.config.js`))
		packagePaths = append(packagePaths, (prefix + `package.json`))
		indexJSPaths = append(indexJSPaths, (prefix + repositoryPathToIndexJS + `index.js`))
	}
	return webpackPaths, packagePaths, indexJSPaths
}

//GenerateCombinedSet generates a new component set based upon the requested sets
func GenerateCombinedSet(c *fiber.Ctx) error {
	requestBody := new(RequestBody)

	if err := c.BodyParser(requestBody); err != nil {
		return err
	}

	if len(requestBody.Organisation) == 0 {
		SendErrorMessage(c, 422, `Please provide a organisation`)
	}

	if len(requestBody.App) == 0 {
		SendErrorMessage(c, 422, `Please provide an organisation`)
	}

	if len(requestBody.Sets) == 0 {
		SendErrorMessage(c, 422, `Please provide an app`)
	}

	outputPrefix := `./public/generated/`
	outputDirectory := outputPrefix + requestBody.Organisation + `/` + requestBody.App
	outputURL := `api/v1/sets/uploaded/` + requestBody.Organisation + `/` + requestBody.App + `/dist`
	webpackPaths, packagePaths, indexJSPaths := getPaths(requestBody)

	generator.CreateDiretoryIfNotExist(outputPrefix + `/` + requestBody.Organisation)
	generator.CreateDiretoryIfNotExist(outputPrefix + `/` + requestBody.Organisation + `/` + requestBody.App)
	err := merger.MergePackages((outputDirectory + `/package.json`), packagePaths)
	if err != nil {
		return err
	}
	merger.MergeIndexJSFiles((outputDirectory + `/index.js`), indexJSPaths)
	merger.MergeWebpackConfigJSs((outputDirectory + `/webpack.config.js`), webpackPaths)

	go generator.SetGenerator(outputDirectory)
	c.Status(200).JSON(&fiber.Map{
		"succes": true,
		"url":    outputURL,
	})
	return nil
}
