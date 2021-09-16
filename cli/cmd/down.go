package cmd

import (
	"github.com/docker/compose/v2/pkg/api"
	"github.com/spf13/cobra"
)

func DownCommand(backend api.Service) *cobra.Command {
	downCmd := &cobra.Command{
		Use:   "down",
		Short: "Stop the dev environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			project, err := convertToProject(cmd.Flag("devfile").Value.String())
			if err != nil {
				return err
			}
			return backend.Down(cmd.Context(), project.Name, api.DownOptions{
				RemoveOrphans: true,
				Project:       project,
			})
		},
	}
	return downCmd
}
