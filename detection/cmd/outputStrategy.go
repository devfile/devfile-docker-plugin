package cmd

import (
	"errors"
	"fmt"
	"github.com/devfile/devrunner/detection/miniBenchmarker"
	"github.com/devfile/library/pkg/devfile/parser"
	devfileCtx "github.com/devfile/library/pkg/devfile/parser/context"
	"github.com/devfile/library/pkg/devfile/parser/data"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"sigs.k8s.io/yaml"
)

const (
	ConsoleOutputFlag   = "console"
	DirectoryOutputFlag = "dir"
)

func GetOutputStrategy(args *OutputStrategyAdditionalArguments) (OutputStrategy, error) {
	var strategy OutputStrategy
	switch args.OutputType {
	case DirectoryOutputFlag:
		strategy = &FileOutputStrategy{}
	case ConsoleOutputFlag:
		strategy = &ConsoleOutputStrategy{}
	default:
		return nil, errors.New("unknown flag")
	}

	err := strategy.ReadAdditionalArgs(args)
	if err != nil {
		return nil, err
	}
	return strategy, nil
}

type OutputStrategyAdditionalArguments struct {
	OutputType    string
	DirectoryPath string
}

type OutputStrategy interface {
	Execute(devfile data.DevfileData) error
	ReadAdditionalArgs(args *OutputStrategyAdditionalArguments) error
}

type ConsoleOutputStrategy struct{}
type FileOutputStrategy struct {
	path string
}

func (*ConsoleOutputStrategy) ReadAdditionalArgs(args *OutputStrategyAdditionalArguments) error {
	return nil
}

func (*ConsoleOutputStrategy) Execute(devfile data.DevfileData) error {
	log.Info("Running ConsoleOutputStrategy")
	yamlData, err := yaml.Marshal(devfile)
	if err != nil {
		return err
	}
	fmt.Println(string(yamlData))

	return nil
}

func (this *FileOutputStrategy) ReadAdditionalArgs(args *OutputStrategyAdditionalArguments) error {
	this.path = args.DirectoryPath

	return nil
}

func (this *FileOutputStrategy) Execute(devfile data.DevfileData) error {
	miniBenchmarker.GetInstance().StartStage("FileOutputStrategy.Output")
	log.Info("Running FileOutputStrategy")
	ctx := devfileCtx.NewDevfileCtx(filepath.Join(this.path, "devfile.yaml"))

	_ = ctx.SetAbsPath()
	d := parser.DevfileObj{
		Ctx:  ctx,
		Data: devfile,
	}

	miniBenchmarker.GetInstance().EndStage("FileOutputStrategy.Output")

	return d.WriteYamlDevfile()
}
