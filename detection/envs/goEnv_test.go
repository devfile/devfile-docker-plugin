package envs

import (
	"github.com/devfile/devrunner/detection/devfileUtils"
	"github.com/devfile/devrunner/detection/util"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"testing"
)

const (
	GoModGoldFile = "go.mod.gold"
)

func TestGoEnvCorrect(t *testing.T) {
	goEnv := goEnv{}

	err := goEnv.TryRespond(util.GetTestDataPath("goEnvCorrect"), GoModGoldFile)
	if err != nil {
		t.Error(err)
		return
	}

	theDevFile, _ := devfileUtils.GetEmptyDevfileData()
	err = goEnv.Build(theDevFile)
	if err != nil {
		t.Error(err)
		return
	}
	components, err := theDevFile.GetComponents(common.DevfileOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	hasGoComponent := false
	for _, component := range components {
		if component.Name == GoComponentName {
			hasGoComponent = true
			break
		}
	}
	if !hasGoComponent {
		t.Error("go component was not found")
		return
	}
	if goEnv.goVersion != "1.16" {
		t.Errorf("got bad go version, got %s, expected %s", goEnv.goVersion, "1.16")
		return
	}
}

func TestGoBadGoMod(t *testing.T) {
	goEnv := goEnv{}

	err := goEnv.TryRespond(util.GetTestDataPath("goEnvBadGoMod"), GoModGoldFile)
	if err == nil {
		t.Error("Did not return an error")
	}
}

func TestGoEmptyGoMod(t *testing.T) {
	goEnv := goEnv{}

	err := goEnv.TryRespond(util.GetTestDataPath("goEnvEmptyGoMod"), GoModGoldFile)
	if err == nil {
		t.Error("Did not return an error")
	}
}

func TestGoEmptyProject(t *testing.T) {
	env := goEnv{}

	err := env.TryRespond(util.GetTestDataPath("emptyProject"), GoModGoldFile)
	if err == nil {
		t.Error("Did not return an error")
	}
}
