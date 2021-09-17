package lib

import (
	"path/filepath"
	"testing"
)

func BaseTest(projectsDir string, projectName string, t *testing.T) {
	thePath := filepath.Join(projectsDir, projectName)
	runner, err := NewDevRunnerTestRunner(thePath)
	if err != nil {
		t.Fatalf("Failed to initialize test runner %s: %s", projectName, err.Error())
		return
	}
	runner.ExecuteComparerWithPath(thePath, t)
}
