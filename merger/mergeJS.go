package merger

import (
	"bufio"
	"fmt"
	"os"
)

// IndexJS is the general struct of a Betty Blocks components set index.js
type IndexJS struct {
	Imports []string
}

// Imports is the type of import statements in js
type Imports []string

// ExportDefaults is the type of export defualt object in js
type ExportDefaults map[string]interface{}

// GetIndexJSFiles gets the indexJS files and appends them to one line by line
func GetIndexJSFiles(files []string) ([]string, []string) {
	var imports []string
	var exports []string
	for _, item := range files {
		file, err := os.Open(item)
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

		jsImports, jsExports := SplitSliceOnValue(fileText, `export default {`)

		for _, item := range jsImports {
			imports = append(imports, item)
		}

		for _, item := range jsExports {
			if item != "export default {" && item != `};` {
				exports = append(exports, item)
			}
		}

	}
	return imports, exports
}

//MergeIndexJSFiles merges index.js files into one index.js file
func MergeIndexJSFiles(newFilePath string, paths []string) {
	newImportText, newExportText := GetIndexJSFiles(paths)
	newImportText = MergeSlice(newImportText)
	newExportText = MergeSlice(newExportText)
	newExportText = append([]string{``, `export default {`}, newExportText...)
	newExportText = append(newExportText, `};`)
	newIndexJSText := ConcatStrings(newImportText, newExportText)
	PrintLines(newFilePath, newIndexJSText)
}
