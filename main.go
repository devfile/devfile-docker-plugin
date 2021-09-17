package main

import (
	"github.com/devfile/devrunner/cli/cmd"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	"github.com/spf13/cobra"
)

func main() {
	plugin.Run(func(dockerCli command.Cli) *cobra.Command {
		var (
			devfilePath, format string
		)

		lazyInit := api.NewServiceProxy()

		up := cmd.UpCommand(lazyInit)
		down := cmd.DownCommand(lazyInit)
		exec := cmd.ExecCommand(lazyInit)
		describe := cmd.DescribeCommand(lazyInit)
		detect := cmd.DetectCommand()

		mainCmd := &cobra.Command{
			Use:   "devenv",
			Short: "A plugin for containerized development environments.",
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				if err := plugin.PersistentPreRunE(cmd, args); err != nil {
					return err
				}
				lazyInit.WithService(compose.NewComposeService(dockerCli.Client(), dockerCli.ConfigFile()))
				return nil
			},
		}

		flags := mainCmd.PersistentFlags()
		flags.StringVar(&devfilePath, "devfile", "devfile.yaml", "The devfile path")
		flags.StringVar(&format, "output", "json", "The output format (e.g. json, text).")

		mainCmd.AddCommand(up, down, exec, describe, detect)
		return mainCmd
	},
		manager.Metadata{
			SchemaVersion: "0.1.0",
			Vendor:        "AWS",
			Version:       "0.0.1-devel",
		})
}
