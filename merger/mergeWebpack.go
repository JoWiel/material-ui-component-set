package merger

import (
	"bufio"
	"fmt"
	"os"
)

//GetWebPacks gets the webpacks from the paths
func GetWebPacks(paths []string) []string {
	var combinedModules []string
	for i, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var fileText []string

		for scanner.Scan() {
			fileText = append(fileText, scanner.Text())
		}
		defer file.Close()
		moduleExport := SplitSliceBetweenValues(fileText, `module.exports = {`, `};`)
		if i == 0 {
			moduleExport = append([]string{`module.exports = [{`}, moduleExport...)
		} else {
			moduleExport = append([]string{`{`}, moduleExport...)
		}
		if i == len(paths)-1 {
			moduleExport = append(moduleExport, `}]`)
		} else {
			moduleExport = append(moduleExport, `},`)
		}
		combinedModules = append(combinedModules, moduleExport...)
	}
	firstLine := `const path = require('path');`
	var newWebPack []string
	newWebPack = append([]string{firstLine, ``}, newWebPack...)
	newWebPack = append(newWebPack, combinedModules...)
	return newWebPack
}

//MergeWebpackConfigJSs merges webpack.config.js in to one for BB Components
func MergeWebpackConfigJSs(newFilePath string, paths []string) {
	newWebpackText := GetWebPacks(paths)
	PrintLines(newFilePath, newWebpackText)
}
