package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	cmd2 "github.com/devfile/devrunner/detection/cmd"
	"github.com/devfile/devrunner/detection/devfileUtils"
	"github.com/devfile/devrunner/detection/envs"
)

type detectOps struct {
	pathType   string
	thePath    string
	outputType string
}

func DetectCommand() *cobra.Command {
	opts := detectOps{}

	detectCmd := &cobra.Command{
		Use: "detect",
		Short: "Analyze the source code and generates an appropriate devfile",
		Run: func(cmd *cobra.Command, args []string) {
			err := validateStringArg("type", opts.pathType, []string{cmd2.DirTypeFlag, cmd2.GitTypeFlag})
			if err != nil {
				log.Fatalf(err.Error())
				return
			}
			err = validateStringArg("outputType", opts.outputType,
				[]string{cmd2.DirectoryOutputFlag, cmd2.ConsoleOutputFlag},
			)
			if err != nil {
				log.Fatalf(err.Error())
				return
			}
			err = executor(opts.pathType, &cmd2.OutputStrategyAdditionalArguments{
				OutputType:    opts.outputType,
				DirectoryPath: opts.thePath,
			})
			if err != nil {
				log.Fatalf(err.Error())
			}
		},
	}

	detectCmd.Flags().StringVarP(&opts.pathType, "type", "t", "dir",
		"type of path [dir, git]",
	)
	detectCmd.Flags().StringVarP(&opts.thePath, "path", "p", ".",
		"path to code",
	)
	detectCmd.Flags().StringVarP(&opts.outputType, "outputType", "o", "console",
		"how to output devfile [dir, console]",
	)

	return detectCmd
}

func validateStringArg(name string, arg string, allowed []string) error {
	for _, s := range allowed {
		if s == arg {
			return nil
		}
	}
	return fmt.Errorf("argument '%s' value not allowed: %s", name, arg)
}

func executor(pathType string, outputArgs *cmd2.OutputStrategyAdditionalArguments) error {
	inputStrategy, err := cmd2.GetInputStrategy(pathType)
	if err != nil {
		return err
	}

	processedPath, err := inputStrategy.GetPath(outputArgs.DirectoryPath)
	if err != nil {
		return err
	}

	theDevFile, err := devfileUtils.GetEmptyDevfileData()

	err = envs.ProcessPath(processedPath, theDevFile)
	if err != nil {
		return err
	}

	outputStr, err := cmd2.GetOutputStrategy(outputArgs)
	if err != nil {
		return err
	}
	err = outputStr.Execute(theDevFile)
	if err != nil {
		return err
	}

	return nil
}
