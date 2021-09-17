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

func TestFlaskExampleAppProject(t *testing.T) { lib.BaseTest(projectsDir, "flaskExampleApp", t) }

func TestPetclinicProject(t *testing.T) { lib.BaseTest(projectsDir, "petclinic", t) }

func TestPetclinicMultimoduleProject(t *testing.T) {
	lib.BaseTest(projectsDir, "petclinicMultimodule", t)
}
