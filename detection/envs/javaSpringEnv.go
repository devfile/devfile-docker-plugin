package envs

import (
	"errors"
	"github.com/devfile/devrunner/detection/envs/additionalComponents"
	"github.com/devfile/devrunner/detection/envs/javaHelpers"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/magiconair/properties"
	"io/fs"
	"os"
	"path/filepath"
)

var MatchFoundError = errors.New("MATCH_FOUND")

const (
	ApplicationPropsFileName = "application.properties"
	ResourcesFolderName      = "resources"
	PropertiesMask           = "*.properties"
	SpringFrameworkBoot      = "org.springframework.boot"
)

var (
	NoSpringError = errors.New("this proj does not have spring")
)

// javaSpringEnv is not a real Env, but an Env which should be used with Java envs
type javaSpringEnv struct {
	hasMySql bool
}

func (*javaSpringEnv) Name() string {
	return "Java (Spring)"
}

func (this *javaSpringEnv) TryRespond(rootPath string, additionalParams ...interface{}) ([]processPathPair, error) {
	if len(additionalParams) == 0 {
		return nil, errors.New("not enough params")
	}

	parentEnv, isOk := additionalParams[0].(javaHelpers.JavaEnvWithDependencies)
	if !isOk {
		return nil, errors.New("not type JavaEnvWithDependencies")
	}

	hasSpring := false
	for _, dep := range parentEnv.GetDependencies().Dependencies {
		if dep.Name == SpringFrameworkBoot {
			hasSpring = true
			break
		}
	}

	if !hasSpring {
		return nil, NoSpringError
	}

	folder, err := this.findResourcesFolder(rootPath)
	if err != nil {
		return nil, err
	}

	this.hasMySql = false

	matches, err := filepath.Glob(filepath.Join(folder, PropertiesMask))
	if err != nil {
		return nil, err
	}

	for _, propertyFile := range matches {
		p := properties.MustLoadFile(propertyFile, properties.UTF8)
		databaseType, isOk := p.Get("database")
		if !isOk {
			continue
		}
		switch databaseType {
		case "mysql":
			this.hasMySql = true
		}
	}

	return nil, nil
}

func (this *javaSpringEnv) Build(devfile data.DevfileData) error {
	err := additionalComponents.BuildMySqlComponent(this.hasMySql, devfile)
	if err != nil {
		return err
	}

	return nil
}

func (*javaSpringEnv) findResourcesFolder(rootPath string) (string, error) {
	var theDir string
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == ResourcesFolderName { // todo: maybe 'resources' folder is always at src/main/resources???
			if _, err := os.Stat(filepath.Join(path, ApplicationPropsFileName)); !os.IsNotExist(err) {
				theDir = path
				return MatchFoundError
			}
		}
		return nil
	})
	if err != nil && err != MatchFoundError {
		return "", err
	}

	if err == nil {
		return "", errors.New("resources dir not found")
	}

	return theDir, nil
}
