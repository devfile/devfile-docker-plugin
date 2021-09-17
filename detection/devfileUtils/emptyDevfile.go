package devfileUtils

import (
	devfilepkg "github.com/devfile/api/v2/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser/data"
)

func GetEmptyDevfileData() (theDevFile data.DevfileData, err error) {
	version := string(data.APISchemaVersion210)
	theDevFile, err = data.NewDevfileData(version)
	theDevFile.SetSchemaVersion(version) // without this line schemaVersion would be "" (empty)
	theDevFile.SetMetadata(devfilepkg.DevfileMetadata{
		Tags: []string{"devrunner-generated"},
	})
	return
}
