package devfile

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

func MaybeMountSources(container v1alpha2.ContainerComponent, service *types.ServiceConfig, projectPath string) {
	mountSources := container.MountSources
	sourcePath := container.SourceMapping
	if sourcePath == "" {
		sourcePath = "/projects"
	}

	if mountSources == nil || (mountSources != nil && *mountSources) {
		service.Volumes = append(service.Volumes, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: projectPath,
			Target: sourcePath,
		})
	}
}
