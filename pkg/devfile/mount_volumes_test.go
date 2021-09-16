package devfile

import (
	"errors"
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	. "github.com/devfile/devrunner/tests"
)

func TestMountVolumes(t *testing.T) {
	t.Run("it adds additional volumes", func(t *testing.T) {
		container := v1alpha2.ContainerComponent{}
		service := types.ServiceConfig{}
		volumes := []string{"/store:/data", "/store2:/data2"}
		err := MountVolumes(container, &service, volumes)

		if !ExpectNil(t, err, "no errors") {
			return
		}

		if !ExpectEqual(t, len(service.Volumes), 2, "wrong number of volumes") {
			return
		}

		ExpectEqual(t, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: "/store",
			Target: "/data",
		}, service.Volumes[0], "first volume")
		ExpectEqual(t, types.ServiceVolumeConfig{
			Type:   "bind",
			Source: "/store2",
			Target: "/data2",
		}, service.Volumes[1], "second volume")
	})

	t.Run("it returns an error if a volume has the wrong format", func(t *testing.T) {
		container := v1alpha2.ContainerComponent{}
		service := types.ServiceConfig{}
		volumes := []string{"/store:/data", "AAAA"}
		err := MountVolumes(container, &service, volumes)

		ExpectedErr := errors.New("invalid format for 'volume' option 'AAAA'. the format should be host_path:container_path")
		ExpectEqual(t, err, ExpectedErr, "Expected error")
	})
}
