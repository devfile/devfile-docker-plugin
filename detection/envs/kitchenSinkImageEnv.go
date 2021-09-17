package envs

import (
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser/data"
)

// No tests required for kitchen sink image.

type kitchenSinkImageEnv struct{}

func (*kitchenSinkImageEnv) GetAdditionalEnvs() []Env {
	return []Env{}
}

func (*kitchenSinkImageEnv) Name() string {
	return "Default image for non-detected or multiple environments"
}

// TryRespond for kitchenSinkImageEnv won't be called.
func (*kitchenSinkImageEnv) TryRespond(rootPath string, additionalParams ...interface{}) error {
	return nil
}

func (*kitchenSinkImageEnv) Build(devfile data.DevfileData) error {
	err := devfile.AddComponents([]v1alpha2.Component{
		{
			Name: "kitchenSink",
			ComponentUnion: v1alpha2.ComponentUnion{
				ComponentType: v1alpha2.ContainerComponentType,
				Container: &v1alpha2.ContainerComponent{
					Container: v1alpha2.Container{
						Image:       "KITCHEN_SINK_TODO",
						MemoryLimit: "7Gb",
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
