package envs

import (
	"errors"
	"fmt"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"github.com/magiconair/properties"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

type pythonEnv struct {
	version string
}

func (this *pythonEnv) Name() string {
	return "Python"
}

func findPythonVersion(path string) (string, error) {
	pipfilePath := filepath.Join(path, "Pipfile")
	if _, err := os.Stat(pipfilePath); err == nil {
		p := properties.MustLoadFile(pipfilePath, properties.UTF8)
		version, isOK := p.Get("python_version")
		if isOK {
			version = strings.ReplaceAll(version, "\"", "")
			return version, nil
		} else {
			log.Warn("did not find python_version")
		}
	}

	reqPath := filepath.Join(path, "Pipfile")
	if _, err := os.Stat(reqPath); err == nil {
		// Will just grab the latest
		return "latest", nil
	}

	return "", errors.New("did not find anything related to Python")
}

func (this *pythonEnv) TryRespond(rootPath string, additionalParams ...interface{}) ([]processPathPair, error) {
	version, err := findPythonVersion(rootPath)
	if err != nil {
		return nil, err
	}

	this.version = version

	return nil, nil
}

func (this *pythonEnv) Build(devfile data.DevfileData) error {
	//TODO implement me
	//panic("implement me")

	err := devfile.AddComponents([]v1alpha2.Component{
		{
			Name: "python",
			ComponentUnion: v1alpha2.ComponentUnion{
				Container: &v1alpha2.ContainerComponent{
					Container: v1alpha2.Container{
						Image:       this.getImageName(),
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

func (this *pythonEnv) getImageName() string {
	if this.version == "latest" {
		return "python:latest"
	}
	return fmt.Sprintf("python:%s-slim", this.version)
}

func (this *pythonEnv) GetAdditionalEnvs() []Env {
	return nil
}
