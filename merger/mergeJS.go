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

// // GetIndexJSFiles2 gets the index.js files
// func GetIndexJSFiles2(files []string) ([]Imports, []ExportDefaults) {
// 	// result := make(map[string]interface{}, len(files))
// 	var imports []Imports
// 	var exportDefaults []ExportDefaults
// 	for i := range files {
// 		file, err := os.Open(files[i])
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		if file == nil {
// 			fmt.Println(`No such file found: ` + files[i])
// 		}
// 		defer file.Close()
// 		jsonBytes, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		var oneFileImports Imports
// 		str := string(jsonBytes)
// 		fmt.Println(str)
// 		// var JSONResult map[string]interface{}
// 		// err = json.Unmarshal([]byte(jsonBytes), &oneFileImports)
// 		err = json.Unmarshal(jsonBytes, &oneFileImports)
// 		imports = append(imports, oneFileImports)
// 	}
// 	return imports, exportDefaults
// }

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

// // GetIndexJSFiles3 gets the indexJS files and appends them to one line by line
// func GetIndexJSFiles3(files []string) ([]Imports, []ExportDefaults) {
// 	// result := make(map[string]interface{}, len(files))
// 	var imports []Imports
// 	var exportDefaults []ExportDefaults
// 	for _, item := range files {
// 		var str strings.Builder
// 		var fileImport Imports
// 		var fileExportDefault ExportDefaults
// 		file, err := os.Open(item)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		scanner := bufio.NewScanner(file)
// 		scanner.Split(bufio.ScanLines)

// 		var fileText []string

// 		for scanner.Scan() {
// 			fileText = append(fileText, scanner.Text())
// 		}
// 		defer file.Close()

// 		fileImport, fileExportDefaultSlice := SplitSliceAfterValue(fileText, `export default {`)
// 		// fileExportDefault = []string{fileExportDefaultSlice}
// 		for i, item := range fileExportDefaultSlice {
// 			if i == 0 {
// 				item = `{ ` + item
// 			}
// 			str.WriteString(item)
// 		}
// 		s := str.String()
// 		fmt.Println(s)
// 		newArr := []byte(s)
// 		err = json.Unmarshal(newArr, &fileExportDefault)

// 		imports = append(imports, fileImport)
// 		exportDefaults = append(exportDefaults, fileExportDefault)
// 		// for _, item := range fileImport {

// 		// }

// 	}
// 	return imports, exportDefaults
// }

// // GenerateIndexJS generates new index.js
// func GenerateIndexJS(imports []string, exports []string) {
// 	var str strings.Builder
// 	for _, item := range imports {
// 		str.WriteString(item)
// 	}
// 	str.WriteString(`export default {`)
// 	for _, item := range exports {
// 		str.WriteString(item)
// 	}
// 	str.WriteString(`};`)
// 	s := str.String()
// 	fmt.Println(s)
// }

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
