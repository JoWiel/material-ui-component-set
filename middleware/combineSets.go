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

func getPaths(requestBody *RequestBody) ([]string, []string, []string, []string) {
	pathToRepositories := `./public/repositories/`
	repositoryPathToIndexJS := `src/`
	var webpackPaths []string
	var packagePaths []string
	var indexJSPaths []string
	var srcPaths []string

	for i := range requestBody.Sets {
		prefix := pathToRepositories + requestBody.Sets[i] + `/`
		webpackPaths = append(webpackPaths, (prefix + `webpack.config.js`))
		packagePaths = append(packagePaths, (prefix + `package.json`))
		indexJSPaths = append(indexJSPaths, (prefix + repositoryPathToIndexJS + `index.js`))
		srcPaths = append(srcPaths, (prefix + `src`))
	}
	return webpackPaths, packagePaths, srcPaths, indexJSPaths
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
	orgAndAppPath := requestBody.Organisation + `/` + requestBody.App
	outputDirectory := outputPrefix + orgAndAppPath
	outputURL := `/api/v1/sets/merged/` + orgAndAppPath
	buildDirectory := `./public/build/` + orgAndAppPath
	webpackPaths, packagePaths, srcPaths, indexJSPaths := getPaths(requestBody)

	generator.CreateDiretoryIfNotExist(outputPrefix + `/` + requestBody.Organisation)
	generator.CreateDiretoryIfNotExist(outputPrefix + `/` + orgAndAppPath)
	generator.CreateDiretoryIfNotExist(outputPrefix + `/` + orgAndAppPath + `/src`)

	merger.CopyDirectory(srcPaths[0], outputPrefix+`/`+orgAndAppPath+`/src`)
	merger.CopyDirectories(srcPaths, outputDirectory)
	err := merger.MergePackages((outputDirectory + `/package.json`), packagePaths)
	if err != nil {
		return err
	}
	merger.MergeIndexJSFiles((outputDirectory + `/src/index.js`), indexJSPaths)
	merger.MergeWebpackConfigJSs((outputDirectory + `/webpack.config.js`), webpackPaths)

	generator.CreateDiretoryIfNotExist(`./public/build/` + requestBody.Organisation)
	generator.CreateDiretoryIfNotExist(buildDirectory)

	//Async to reduce timeout on request side
	go func() {
		err = generator.BuildSet(outputDirectory, buildDirectory)
	}()

	if err != nil {
		return err
	}
	c.Status(200).JSON(&fiber.Map{
		"succes":  true,
		"message": "Component-set is being build this will take approximately 5 minutes.",
		"url":     outputURL,
	})
	return nil
}
