package devfile

import (
	"fmt"
	"strings"

	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

func MountVolumes(container v1alpha2.ContainerComponent, service *types.ServiceConfig, volumes []string) error {
	for _, volume := range volumes {
		parts := strings.Split(volume, ":")
		if len(parts) != 2 {
			return fmt.Errorf("invalid format for 'volume' option '%s'. the format should be host_path:container_path", volume)
		}
		service.Volumes = append(service.Volumes, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: parts[0],
			Target: parts[1],
		})
	}
	return nil
}
