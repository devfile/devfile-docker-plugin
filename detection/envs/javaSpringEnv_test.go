package envs

import (
	"github.com/devfile/devrunner/detection/devfileUtils"
	"github.com/devfile/devrunner/detection/envs/additionalComponents"
	"github.com/devfile/devrunner/detection/envs/javaHelpers"
	"github.com/devfile/devrunner/detection/util"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"testing"
)

type testJavaEnv struct {
	Dependencies javaHelpers.JavaDependencies
}

func (this testJavaEnv) GetDependencies() javaHelpers.JavaDependencies {
	return this.Dependencies
}

var correctTestJavaEnv = testJavaEnv{
	Dependencies: javaHelpers.JavaDependencies{
		Dependencies: []javaHelpers.JavaDependenciesPair{
			{
				Name: SpringFrameworkBoot,
			},
		},
	},
}

func TestJavaSpringEnvDependencyResolverGood(t *testing.T) {
	springEnv := javaSpringEnv{}
	_, err := springEnv.TryRespond("", correctTestJavaEnv)
	if err == NoSpringError {
		t.Errorf("expected %s, got %s", NoSpringError.Error(), err.Error())
	}
}

func TestJavaSpringEnvDependencyResolverBad(t *testing.T) {
	testJavaEnv := testJavaEnv{
		Dependencies: javaHelpers.JavaDependencies{
			Dependencies: []javaHelpers.JavaDependenciesPair{
				{
					Name: "NeverGonnaGiveYouUp",
				},
			},
		},
	}

	springEnv := javaSpringEnv{}
	_, err := springEnv.TryRespond("", testJavaEnv)
	if err != NoSpringError {
		t.Errorf("expected %s, got %s", NoSpringError.Error(), err.Error())
	}
}

func TestJavaSpringEnvCorrect(t *testing.T) {
	springEnv := javaSpringEnv{}
	_, err := springEnv.TryRespond(util.GetTestDataPath("javaSpringWithMySql"), correctTestJavaEnv)
	if err != nil {
		t.Error(err)
	}

	theDevFile, _ := devfileUtils.GetEmptyDevfileData()
	err = springEnv.Build(theDevFile)
	if err != nil {
		t.Error(err)
	}

	components, err := theDevFile.GetComponents(common.DevfileOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	hasComponent := false
	for _, component := range components {
		if component.Name == additionalComponents.MySqlComponentName {
			hasComponent = true
			break
		}
	}
	if !hasComponent {
		t.Error("mysql component was not found")
		return
	}
}
