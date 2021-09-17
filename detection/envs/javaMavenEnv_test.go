package envs

import (
	"devrunner/devfileUtils"
	"devrunner/util"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"testing"
)

func TestJavaMavenEnv(t *testing.T) {
	javaMavenEnv := javaMavenEnv{}

	_, err := javaMavenEnv.TryRespond(util.GetTestDataPath("javaMavenCorrect"), GoModGoldFile)
	if err != nil {
		t.Error(err)
		return
	}

	theDevFile, _ := devfileUtils.GetEmptyDevfileData()
	err = javaMavenEnv.Build(theDevFile)
	if err != nil {
		t.Error(err)
		return
	}
	components, err := theDevFile.GetComponents(common.DevfileOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	hasComponent := false
	for _, component := range components {
		if component.Name == MavenComponentName {
			hasComponent = true
			break
		}
	}
	if !hasComponent {
		t.Error("go component was not found")
		return
	}
	if javaMavenEnv.javaVersion != "8" {
		t.Errorf("got bad javaHelpers version, got %s, expected %s", javaMavenEnv.javaVersion, "8")
		return
	}
}

func TestMavenNoVersion(t *testing.T) {
	env := javaMavenEnv{}

	_, err := env.TryRespond(util.GetTestDataPath("javaMavenNoVersion"), GoModGoldFile)
	if err == nil {
		t.Error("Did not return an error")
	}
}

func TestMavenEmptyProject(t *testing.T) {
	env := javaMavenEnv{}

	_, err := env.TryRespond(util.GetTestDataPath("emptyProject"), GoModGoldFile)
	if err == nil {
		t.Error("Did not return an error")
	}
}
