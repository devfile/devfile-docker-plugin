package comparers

import (
	"fmt"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
)

type ContainerComparer struct{}

func (*ContainerComparer) Name() string {
	return "Container"
}

func (*ContainerComparer) Compare(iDevfile data.DevfileData, jDevfile data.DevfileData) error {
	iComponents, err := iDevfile.GetComponents(common.DevfileOptions{})
	if err != nil {
		return err
	}
	jComponents, err := jDevfile.GetComponents(common.DevfileOptions{})

	for _, iComponent := range iComponents {
		foundNeedle := false
		for _, jComponent := range jComponents {
			if iComponent.Name != jComponent.Name {
				continue
			}
			if iComponent.Container.Image != jComponent.Container.Image {
				continue
			}
			foundNeedle = true
		}
		if !foundNeedle {
			return fmt.Errorf("%s was not found or different in jComponent", iComponent.Name)
		}
	}

	return nil
}
