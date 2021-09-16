package cmd

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/docker/compose/v2/pkg/api"
)

type svcDescription struct {
	Service string
	Status  string
}

func DescribeCommand(backend api.Service) *cobra.Command {
	describeCmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe a dev environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			devfile, err := cmd.Flags().GetString("devfile")
			if err != nil {
				return err
			}
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}
			project, err := convertToProject(devfile)
			if err != nil {
				return err
			}

			containers, err := backend.Ps(cmd.Context(), project.Name, api.PsOptions{
				All: true,
			})
			if err != nil {
				return err
			}

			var services []svcDescription
			for _, cnt := range containers {
				services = append(services, svcDescription{
					Service: cnt.Service,
					Status:  cnt.State,
				})
			}

			if output == "json" {
				return jsonOutput(cmd.OutOrStdout(), services)
			}

			return textOutput(cmd.OutOrStdout(), services)
		},
	}
	return describeCmd
}

func textOutput(writer io.Writer, services []svcDescription) error {
	fmt.Fprintf(writer, "Name\tStatus\n")
	for _, svc := range services {
		fmt.Fprintf(writer, "%s\t%s\n", svc.Service, svc.Status)
	}

	fmt.Fprintln(writer)

	return nil
}

func jsonOutput(writer io.Writer, services []svcDescription) error {
	out, err := json.MarshalIndent(services, "", " ")
	if err != nil {
		return err
	}
	_, err = writer.Write(out)
	return err
}
