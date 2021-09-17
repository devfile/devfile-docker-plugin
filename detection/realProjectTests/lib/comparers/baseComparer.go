package comparers

import (
	"github.com/devfile/library/pkg/devfile/parser/data"
)

func GetComparers() []BaseComparer {
	return []BaseComparer{
		&VersionComparer{},
		&ContainerComparer{},
	}
}

type BaseComparer interface {
	Name() string
	Compare(iDevfile data.DevfileData, jDevfile data.DevfileData) error
}
