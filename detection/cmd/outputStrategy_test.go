package cmd

import (
	"github.com/devfile/devrunner/detection/devfileUtils"
	"github.com/devfile/devrunner/detection/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	ExceptedOutput = "metadata:\n  tags:\n  - devrunner-generated\nschemaVersion: 2.1.0\n\n"
)

func TestGetOutputStrategy(t *testing.T) {
	dirOutputStrategy, err := GetOutputStrategy(&OutputStrategyAdditionalArguments{
		OutputType: DirectoryOutputFlag,
	})
	if err != nil {
		t.Error(err)
	}
	_, isOk := dirOutputStrategy.(*FileOutputStrategy)
	if !isOk {
		t.Error("Got wrong class for dir output strategy")
	}

	consoleOutputStrategy, err := GetOutputStrategy(&OutputStrategyAdditionalArguments{
		OutputType: ConsoleOutputFlag,
	})
	if err != nil {
		t.Error(err)
	}
	_, isOk = consoleOutputStrategy.(*ConsoleOutputStrategy)
	if !isOk {
		t.Error("Got wrong class for console output strategy")
	}
}

func TestConsoleOutputStrategy(t *testing.T) {
	outputStrategy := ConsoleOutputStrategy{}
	theDevFile, _ := devfileUtils.GetEmptyDevfileData()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := outputStrategy.Execute(theDevFile)
	if err != nil {
		w.Close()
		os.Stdout = rescueStdout
		t.Error(err)
	}

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	outStr := string(out)

	if util.RemoveNewLines(outStr) != util.RemoveNewLines(ExceptedOutput) {
		t.Errorf("got wrong output.\nexpected: \n%s\ngot:\n%s", ExceptedOutput, outStr)
	}
}

func TestFileOutputStrategy(t *testing.T) {
	theDir := os.TempDir()
	outputStrategy := FileOutputStrategy{
		path: theDir,
	}
	theDevFile, _ := devfileUtils.GetEmptyDevfileData()

	err := outputStrategy.Execute(theDevFile)
	if err != nil {
		t.Error(err)
	}

	result, err := ioutil.ReadFile(filepath.Join(theDir, "devfile.yaml"))
	if err != nil {
		t.Error(err)
	}

	resultStr := string(result)

	if util.RemoveNewLines(resultStr) != util.RemoveNewLines(ExceptedOutput) {
		t.Errorf("got wrong output.\nexpected: \n%s\ngot:\n%s", ExceptedOutput, resultStr)
	}
}
