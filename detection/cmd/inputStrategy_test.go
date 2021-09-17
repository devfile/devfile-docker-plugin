package cmd

import (
	"github.com/devfile/devrunner/detection/util"
	"os"
	"strings"
	"testing"
)

func TestGetInputStrategy(t *testing.T) {
	dirInputStrategy, err := GetInputStrategy(DirTypeFlag)
	if err != nil {
		t.Error(err)
	}
	_, isOk := dirInputStrategy.(DirInputStrategy)
	if !isOk {
		t.Error("Got wrong class for dir input strategy")
	}

	gitInputStrategy, err := GetInputStrategy(GitTypeFlag)
	if err != nil {
		t.Error(err)
	}
	_, isOk = gitInputStrategy.(GitInputStrategy)
	if !isOk {
		t.Error("Got wrong class for git input strategy")
	}
}

func TestGitInputStrategy(t *testing.T) {
	t.Skip("Skip")
	path, err := GitInputStrategy{}.GetPath(util.GetTestDataPath("gitRepo"))

	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(path, os.TempDir()) {
		t.Fatalf("Path does not start with temp dir")
	}
}
