package cmd

import (
	"github.com/spf13/cobra"
)

func createTestDevEnvCMD(subCmd *cobra.Command, testDevFile string) *cobra.Command {
	mainCmd := &cobra.Command{
		Use:   "devenv",
		Short: "test wrapper",
	}

	devfilePath := ""
	format := ""
	flags := mainCmd.PersistentFlags()
	flags.StringVar(&devfilePath, "devfile", testDevFile, "The devfile path")
	flags.StringVar(&format, "output", "json", "The output format (e.g. json, text).")

	mainCmd.AddCommand(subCmd)

	return mainCmd
}
