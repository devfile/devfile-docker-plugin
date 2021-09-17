package envs

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/devrunner/detection/devfileUtils"
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
	path         string
	javaVersion  string
	projectName  string
	dependencies []javaHelpers.JavaDependenciesPair
}

func (*javaMavenEnv) GetNewInstance() Env {
	return &javaMavenEnv{}
}

func (*javaMavenEnv) GetAdditionalEnvs() []Env {
	return []Env{
		&javaSpringEnv{},
	}
}

func (*javaMavenEnv) Name() string {
	return "Java (Maven)"
}

func (this *javaMavenEnv) TryRespond(rootPath string, additionalParams ...interface{}) ([]processPathPair, error) {
	project, err := this.buildPom(rootPath)
	if err != nil {
		return nil, err
	}

	additionalPaths := []processPathPair{}
	for _, module := range project.Modules.Module {
		log.Infof("Found additional module %s", module)
		additionalPaths = append(additionalPaths, processPathPair{
			filepath.Join(rootPath, module), &javaMavenEnv{},
		})
	}

	for _, dep := range project.Dependencies.Dependency {
		this.dependencies = append(this.dependencies, javaHelpers.JavaDependenciesPair{
			Name: dep.GroupId,
		})
	}

	if project.Properties.JavaVersion == "" {
		return additionalPaths, errors.New("no java version")
	}

	this.javaVersion, err = javaHelpers.ConvertJavaVersion(project.Properties.JavaVersion)
	if err != nil {
		return additionalPaths, err
	}

	this.projectName = project.Name

	return additionalPaths, nil
}

func (this *javaMavenEnv) Build(devfile data.DevfileData) error {
	if this.projectName != "" {
		currentMetadata := devfile.GetMetadata()
		currentMetadata.Name = this.projectName
		devfile.SetMetadata(currentMetadata)
	}
	if devfileUtils.DevfileHasComponent(devfile, MavenComponentName) {
		return nil
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

func (this *javaMavenEnv) GetDependencies() javaHelpers.JavaDependencies {
	return javaHelpers.JavaDependencies{
		Dependencies: this.dependencies,
	}
}

func (this *javaMavenEnv) buildPom(path string) (*javaHelpers.MavenProject, error) {
	pomFilePath := filepath.Join(path, PomXmlFileName)

	f, err := os.Open(pomFilePath)

	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Warnf("failed to close pom.xml %s", err.Error())
		}
	}(f)

	byteValue, _ := ioutil.ReadAll(f)

	var project javaHelpers.MavenProject
	err = xml.Unmarshal(byteValue, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (this *javaMavenEnv) getMavenDockerImage() string {
	return fmt.Sprintf("maven:3.8-jdk-%s", this.javaVersion) // todo maven-wrapper.properties has specific version
}
