package cmd

import (
	"github.com/devfile/devrunner/detection/util"
	"os"
	"path/filepath"
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
	gitRepoPath := util.GetTestDataPath("testRepo")
	err := os.Rename(filepath.Join(gitRepoPath, ".git_escaped"), filepath.Join(gitRepoPath, ".git"))
	if err != nil {
		t.Fatal(err)
	}
	path, err := GitInputStrategy{}.GetPath(gitRepoPath)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Rename(filepath.Join(gitRepoPath, ".git"), filepath.Join(gitRepoPath, ".git_escaped"))
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(path, os.TempDir()) {
		t.Fatalf("Path does not start with temp dir")
	}
}
