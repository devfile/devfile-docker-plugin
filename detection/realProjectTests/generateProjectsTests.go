package main

import (
	"fmt"
	"github.com/devfile/devrunner/detection/realProjectTests/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Templates
const (
	GeneratedFileHeader = `
// This file was generated automatically and it is not a good idea to edit it.
// Run "Generate realProjectTests" to update it.
package main

import (
	"github.com/devfile/devrunner/detection/realProjectTests/lib"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var baseDir string
var projectsDir = filepath.Join(baseDir, lib.ProjectsDir)

func init() {
	baseDirCurrent, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	baseDirSplit := strings.Split(baseDirCurrent, string(os.PathSeparator))
	if baseDirSplit[len(baseDirSplit)-1] == lib.TestsDir {
		baseDir = baseDirCurrent
	} else if baseDirSplit[len(baseDirSplit)-1] == lib.DevRunnerDir {
		baseDir = filepath.Join(baseDirCurrent, lib.TestsDir)
	} else {
		panic("bad pwd dir")
	}
}
`
	FunctionContents = `
func Test%sProject(t *testing.T)  { lib.BaseTest(projectsDir, "%s", t) }
`
)

func main() {
	baseDevRunnerDir, err := os.Getwd()
	generatedTestFullPath := filepath.Join(baseDevRunnerDir, lib.TopLevelDir, lib.TestsDir, lib.GeneratedTestsFile)
	projectsFullPath := filepath.Join(baseDevRunnerDir, lib.TopLevelDir, lib.TestsDir, lib.ProjectsDir)

	if err != nil {
		println(err.Error())
		return
	}
	if _, err := os.Stat(generatedTestFullPath); err == nil {
		if err := os.Remove(generatedTestFullPath); err != nil {
			println(err.Error())
			return
		}
	}

	files, err := ioutil.ReadDir(projectsFullPath)
	if err != nil {
		println(err.Error())
		return
	}
	var b strings.Builder
	b.Grow(16)
	b.WriteString(GeneratedFileHeader)
	for _, f := range files {
		if f.IsDir() {
			b.WriteString(fmt.Sprintf(FunctionContents, strings.Title(f.Name()), f.Name()))
		}
	}

	err = os.WriteFile(generatedTestFullPath, []byte(b.String()), 0644)
	if err != nil {
		println(err.Error())
		return
	}

	println(b.String())
	println("Done")
}
