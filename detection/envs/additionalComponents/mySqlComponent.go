package additionalComponents

import (
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser/data"
)

const MySqlComponentName = "mysql"

func BuildMySqlComponent(needsToAdd bool, devfile data.DevfileData) error {
	if !needsToAdd {
		return nil
	}

	err := devfile.AddComponents([]v1alpha2.Component{
		{
			Name: MySqlComponentName,
			ComponentUnion: v1alpha2.ComponentUnion{
				Container: &v1alpha2.ContainerComponent{
					Container: v1alpha2.Container{
						Image:       "mysql",
						MemoryLimit: "3Gb",
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
