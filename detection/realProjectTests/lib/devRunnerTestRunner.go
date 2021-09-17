package lib

import (
	"github.com/devfile/devrunner/detection/devfileUtils"
	"github.com/devfile/devrunner/detection/envs"
	"github.com/devfile/devrunner/detection/realProjectTests/lib/comparers"
	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"path/filepath"
	"testing"
)

const (
	DevfileName     = "devfile.yaml"
	DevfileGoldName = DevfileName
)

type DevRunnerTestRunner interface {
	ExecuteComparer(generatedFileData data.DevfileData, t *testing.T)
	ExecuteComparerWithPath(filepath string, t *testing.T)
}

type DevRunnerTestContext struct {
	ProjectPath string

	GoldFileObj data.DevfileData
}

func (this *DevRunnerTestContext) ExecuteComparerWithPath(filepath string, t *testing.T) {
	theDevFile, err := devfileUtils.GetEmptyDevfileData()
	if err != nil {
		t.Fatalf("Failed to initialize default devfile: %s", err.Error())
		return
	}

	if err = envs.ProcessPath(filepath, theDevFile); err != nil {
		t.Fatalf("Failed to process devfile: %s", err.Error())
		return
	}

	this.ExecuteComparer(theDevFile, t)
}

func (this *DevRunnerTestContext) ExecuteComparer(generatedFileData data.DevfileData, t *testing.T) {
	for _, comparer := range comparers.GetComparers() {
		err := comparer.Compare(generatedFileData, this.GoldFileObj)
		if err != nil {
			t.Fatalf("Comparison failed: %s", err.Error())
		}
	}
}

func NewDevRunnerTestRunner(projectPath string) (DevRunnerTestRunner, error) {
	ctx := DevRunnerTestContext{
		ProjectPath: projectPath,
	}

	err := ctx.getGoldFile()
	if err != nil {
		return nil, err
	}

	return &ctx, nil
}

func (this *DevRunnerTestContext) getGoldFile() error {
	devfilePath := filepath.Join(this.ProjectPath, DevfileGoldName)
	devfile, err := parser.ParseDevfile(parser.ParserArgs{
		Path: devfilePath,
	})
	if err != nil {
		return err
	}

	this.GoldFileObj = devfile.Data

	return nil
}
