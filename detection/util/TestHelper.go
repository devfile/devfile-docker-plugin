package util

import (
	"path/filepath"
	"runtime"
	"strings"
)

func GetTestDataPath(myPath string) string {
	_, filename, _, _ := runtime.Caller(1)
	filenameSplit := strings.Split(filename, string(filepath.Separator))
	filenameSplit = append(filenameSplit[:len(filenameSplit)-1], "testData", myPath)

	return strings.Join(filenameSplit, string(filepath.Separator))
}

func RemoveNewLines(str string) string {
	return strings.Replace(str, "\n", "", -1)
}
