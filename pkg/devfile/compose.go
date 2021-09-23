package devfile

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
)

// ToComposeProject converts a devfile object to its corresponding Compose object.
// Currently only container and volume components are supported.
func ToComposeProject(devfile parser.DevfileObj) (types.Project, error) {
	result := types.Project{
		Name:    devfile.Data.GetMetadata().Name,
		Volumes: make(types.Volumes),
	}

	cnts, err := devfile.Data.GetDevfileContainerComponents(common.DevfileOptions{})
	if err != nil {
		return types.Project{}, err
	}

	for _, cnt := range cnts {
		result.Services = append(result.Services, convertToService(cnt))
	}

	vols, err := devfile.Data.GetDevfileVolumeComponents(common.DevfileOptions{})
	if err != nil {
		return types.Project{}, err
	}

	for _, vol := range vols {
		result.Volumes[vol.Name] = convertToVolume(vol)
	}

	return result, nil
}

func convertToService(devContainer v1alpha2.Component) types.ServiceConfig {
	cnt := devContainer.Container

	// TODO: tmp to make devfile from registry to work
	if cnt.Command == nil {
		cnt.Command = []string{"sleep", "infinity"}
	}

	svc := types.ServiceConfig{
		Name:        devContainer.Name,
		Entrypoint:  cnt.Command,
		Command:     cnt.Args,
		Image:       cnt.Image,
		NetworkMode: "host",
		Environment: make(types.MappingWithEquals),
	}

	//TODO: https://i.amazon.com/issues/TIDE-1364 convert other fields like CPU and memory
	for _, kv := range cnt.Env {
		svc.Environment[kv.Name] = &kv.Value
	}

	for _, vols := range cnt.VolumeMounts {
		svc.Volumes = append(svc.Volumes, types.ServiceVolumeConfig{
			Type:   "volume", // TODO: https://i.amazon.com/issues/TIDE-1365 use only bind mounts for volumes so we control migration
			Source: vols.Name,
			Target: vols.Path,
		})
	}

	return svc
}

func convertToVolume(vol v1alpha2.Component) types.VolumeConfig {
	return types.VolumeConfig{
		Name: vol.Name,
	}
}
