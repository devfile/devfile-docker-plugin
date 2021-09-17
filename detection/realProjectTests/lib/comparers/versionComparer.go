package comparers

import (
	"fmt"
	"github.com/devfile/library/pkg/devfile/parser/data"
)

type VersionComparer struct{}

func (*VersionComparer) Name() string {
	return "Version"
}

func (*VersionComparer) Compare(iDevfile data.DevfileData, jDevfile data.DevfileData) error {
	if iDevfile.GetSchemaVersion() != iDevfile.GetSchemaVersion() {
		return fmt.Errorf("versions are different (%s vs %s)", iDevfile.GetSchemaVersion(), jDevfile.GetSchemaVersion())
	}

	return nil
}
