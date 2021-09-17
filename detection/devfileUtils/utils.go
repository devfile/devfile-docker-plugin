package devfileUtils

import (
	devfilepkg "github.com/devfile/api/v2/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	log "github.com/sirupsen/logrus"
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

// DevfileHasComponent is for Build funcs that have multiple modules. This is a dirty hack, and we MUST remove when there will be a proper solution from devfile library.
func DevfileHasComponent(devfile data.DevfileData, name string) bool {
	comps, err := devfile.GetComponents(common.DevfileOptions{})
	if err != nil {
		log.Warnf("failed to read components: %s", err.Error())
		return false
	}
	for _, comp := range comps {
		if comp.Name == name {
			log.Infof("Found component %s", name)
			return true
		}
	}
	return false
}
