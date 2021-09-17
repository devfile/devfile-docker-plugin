package envs

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/devrunner/detection/envs/javaHelpers"
	"github.com/devfile/library/pkg/devfile/parser/data"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	PomXmlFileName     = "pom.xml"
	MavenComponentName = "maven"
)

type javaMavenEnv struct {
	javaVersion  string
	projectName  string
	dependencies []string
}

func (*javaMavenEnv) GetAdditionalEnvs() []Env {
	return []Env{
		&javaSpringEnv{},
	}
}

func (*javaMavenEnv) Name() string {
	return "Java (Maven)"
}

func (this *javaMavenEnv) TryRespond(rootPath string, additionalParams ...interface{}) error {
	pomFilePath := filepath.Join(rootPath, PomXmlFileName)

	f, err := os.Open(pomFilePath)

	if err != nil {
		return err
	} else {
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Warnf("failed to close pom.xml %s", err.Error())
			}
		}(f)
	}

	byteValue, _ := ioutil.ReadAll(f)

	var project javaHelpers.MavenProject
	err = xml.Unmarshal(byteValue, &project)
	if err != nil {
		return err
	}

	if project.Properties.JavaVersion == "" {
		return errors.New("no java version")
	}

	this.javaVersion, err = javaHelpers.ConvertJavaVersion(project.Properties.JavaVersion)
	if err != nil {
		return err
	}

	this.dependencies = []string{}

	for _, dep := range project.Dependencies.Dependency {
		this.dependencies = append(this.dependencies, dep.GroupId)
	}

	this.projectName = project.Name

	return nil
}

func (this *javaMavenEnv) Build(devfile data.DevfileData) error {
	if this.projectName != "" {
		currentMetadata := devfile.GetMetadata()
		currentMetadata.Name = this.projectName
		devfile.SetMetadata(currentMetadata)
	}
	err := devfile.AddComponents([]v1alpha2.Component{
		{
			Name: MavenComponentName,
			ComponentUnion: v1alpha2.ComponentUnion{
				Container: &v1alpha2.ContainerComponent{
					Container: v1alpha2.Container{
						Image:       this.getMavenDockerImage(),
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

func (this *javaMavenEnv) getMavenDockerImage() string {
	return fmt.Sprintf("maven:3.8-jdk-%s", this.javaVersion)
}

func (this *javaMavenEnv) GetDependencies() javaHelpers.JavaDependencies {
	return javaHelpers.JavaDependencies{
		Dependencies: this.dependencies,
	}
}
