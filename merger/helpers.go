package merger

import (
	"github.com/mcuadros/go-version"

	"fmt"
	"os"
)

// AppendIfMissing all the unique values in the slice and return the appended slice
func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// MergeSlice merges als the unique values in the slice and return the appended slice
func MergeSlice(slice []string) []string {
	results := make([]string, 0) // slice tostore the result
	for i := range slice {
		results = AppendIfMissing(results, slice[i])
	}
	return results
}

//ConcatStrings concats two slices of strings
func ConcatStrings(slice1 []string, slice2 []string) []string {
	for _, item := range slice2 {
		slice1 = append(slice1, item)
	}
	return slice1
}

// SplitSliceOnValue splits the slice on a value and return
func SplitSliceOnValue(slice []string, value string) ([]string, []string) {
	for i, item := range slice {
		if item == value {
			return slice[0:i], slice[i:]
		}
	}
	return nil, nil
}

// SplitSliceAfterValue splits the slice on a value and return
func SplitSliceAfterValue(slice []string, value string) ([]string, []string) {
	for i, item := range slice {
		if item == value {
			splitIndex := i + 1
			return slice[0:splitIndex], slice[splitIndex:]
		}
	}
	return nil, nil
}

// MergeMaps merges the unique keys in the maps in to one
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			current := result[k]
			if current == nil || version.CompareSimple(fmt.Sprintf("%v", current), fmt.Sprintf("%v", v)) > 1 {
				result[k] = v
			}

		}
	}
	return result
}

//PrintLines writes strings in to one file
func PrintLines(filePath string, values []string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

//SplitSliceBetweenValues splits the slice between the two given values
func SplitSliceBetweenValues(webpackContent []string, startValue string, endValue string) []string {
	var startIndex int
	var endIndex int
	for i, line := range webpackContent {
		if line == startValue {
			startIndex = i + 1
		}
		if line == endValue {
			endIndex = i
		}
	}
	return webpackContent[startIndex:endIndex]
}
