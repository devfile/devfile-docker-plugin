package cmd

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"github.com/spf13/cobra"

	convert "github.com/devfile/devrunner/pkg/devfile"
	"github.com/docker/compose/v2/pkg/api"
)

type upOps struct {
	environment []string
	volumes     []string
	projectPath string
	workingDir  string
}

func UpCommand(backend api.Service) *cobra.Command {
	opts := upOps{}

	upCmd := &cobra.Command{
		Use:   "up",
		Short: "Create and start dev environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			devFilepath, err := cmd.Flags().GetString("devfile")
			if err != nil {
				return err
			}
			devFile, _, err := devfile.ParseDevfileAndValidate(parser.ParserArgs{
				Path: devFilepath,
			})
			if err != nil {
				return err
			}

			project, err := convert.ToComposeProject(devFile)
			if err != nil {
				return err
			}

			project.WorkingDir = opts.workingDir

			services, err := addAdditionalOpts(devFile, project.Services, opts)
			if err != nil {
				return err
			}
			project.Services = services

			return backend.Up(cmd.Context(), &project, api.UpOptions{
				Create: api.CreateOptions{
					RemoveOrphans: true,
					Inherit:       true,
				},
				Start: api.StartOptions{
					CascadeStop: false,
				},
			})
		},
	}

	upCmd.Flags().StringArrayVarP(&opts.environment, "env", "e", []string{}, "Set environment variables.")
	upCmd.Flags().StringArrayVarP(&opts.volumes, "volume", "v", []string{}, "Mount volumes. The format should be host_path:container_path.")
	upCmd.Flags().StringVarP(&opts.projectPath, "projectpath", "p", ".", "The project path.")
	upCmd.Flags().StringVarP(&opts.workingDir, "workdir", "w", ".", "The working directory path.")

	return upCmd
}

func addAdditionalOpts(devFile parser.DevfileObj, currentServices types.Services, opts upOps) (types.Services, error) {
	components, err := devFile.Data.GetDevfileContainerComponents(common.DevfileOptions{})
	if err != nil {
		return nil, err
	}

	services := make(types.Services, len(currentServices))
	for idx, service := range currentServices {
		service.Environment = service.Environment.OverrideBy(types.NewMappingWithEquals(opts.environment))

		container := components[idx].Container

		convert.MaybeMountSources(*container, &service, opts.projectPath)
		err = convert.MountVolumes(*container, &service, opts.volumes)
		if err != nil {
			return nil, err
		}

		services[idx] = service
	}
	return services, nil
}
