package devfile

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	. "github.com/devfile/devrunner/tests"
)

func TestMaybeMountSources(t *testing.T) {
	type testCase struct {
		Name          string
		Expectation   string
		MountSources  *bool
		SourceMapping string
		ProjectPath   string
		Want          types.ServiceConfig
	}

	tests := []testCase{
		{
			Name:         "MountSources = false",
			Expectation:  "no volume is mounted",
			ProjectPath:  ".",
			MountSources: ToBoolPtr(false),
			Want:         types.ServiceConfig{},
		},
		{
			Name:         "MountSources = true",
			Expectation:  "volume is mounted to default /projects",
			ProjectPath:  ".",
			MountSources: ToBoolPtr(true),
			Want: types.ServiceConfig{
				Volumes: []types.ServiceVolumeConfig{
					{
						Type:   "bind",
						Source: ".",
						Target: "/projects",
					},
				},
			},
		},
		{
			Name:         "MountSources = nil",
			Expectation:  "volume is mounted to default /projects",
			ProjectPath:  ".",
			MountSources: nil,
			Want: types.ServiceConfig{
				Volumes: []types.ServiceVolumeConfig{
					{
						Type:   "bind",
						Source: ".",
						Target: "/projects",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			container := v1alpha2.ContainerComponent{}
			container.MountSources = tt.MountSources
			container.SourceMapping = tt.SourceMapping
			service := types.ServiceConfig{}
			MaybeMountSources(container, &service, tt.ProjectPath)
			ExpectEqual(t, service, tt.Want, tt.Expectation)
		})
	}
}
