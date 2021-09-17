package envs

import (
	"errors"
	"fmt"
	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/pkg/devfile/parser/data"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"path/filepath"
)

const (
	GoModFileName   = "go.mod"
	GoComponentName = "go"
)

type goEnv struct {
	goVersion string
}

func (this *goEnv) GetAdditionalEnvs() []Env {
	return []Env{}
}

func (*goEnv) Name() string {
	return "Go"
}

func (this *goEnv) TryRespond(rootPath string, additionalParams ...interface{}) ([]processPathPair, error) {
	// Override go.mod file name
	goModCurrentFileName := GoModFileName
	if len(additionalParams) == 1 {
		arg1Str, isOk := additionalParams[0].(string)
		if isOk {
			goModCurrentFileName = arg1Str
		}
	}
	goModFilePath := filepath.Join(rootPath, goModCurrentFileName)
	goModFile, err := ioutil.ReadFile(goModFilePath)
	if err != nil {
		return nil, err
	}
	lax, err := modfile.ParseLax(goModFilePath, goModFile, nil)
	if err != nil {
		return nil, err
	}

	if lax.Go == nil {
		return nil, errors.New("got nil for lax.Go")
	}

	this.goVersion = lax.Go.Version

	if this.goVersion == "" {
		return nil, errors.New("got empty string for go version")
	}

	return nil, nil
}

func (this *goEnv) Build(devfile data.DevfileData) error {
	err := devfile.AddComponents([]v1alpha2.Component{
		{
			Name: GoComponentName,
			ComponentUnion: v1alpha2.ComponentUnion{
				ComponentType: v1alpha2.ContainerComponentType,
				Container: &v1alpha2.ContainerComponent{
					Container: v1alpha2.Container{
						Image:       fmt.Sprintf("golang:%s-alpine", this.goVersion),
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
