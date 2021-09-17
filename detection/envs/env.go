package envs

import (
	"github.com/devfile/devrunner/detection/miniBenchmarker"
	"github.com/devfile/library/pkg/devfile/parser/data"
	log "github.com/sirupsen/logrus"
)

func GetEnvs() []TopLevelEnv {
	return []TopLevelEnv{
		&javaMavenEnv{},
		&goEnv{},
		&pythonEnv{},
		// todo gradle
	}
}

func GetDefaultEnv() TopLevelEnv {
	return &kitchenSinkImageEnv{}
}

// ProcessPath is tested at realProjectTests

type processPathPair struct {
	Path string
	Env  TopLevelEnv
}

type Env interface {
	Name() string
	// TryRespond returns []string for paths that should be also processed
	TryRespond(rootPath string, additionalParams ...interface{}) ([]processPathPair, error)
	Build(devfile data.DevfileData) error
}

type TopLevelEnv interface {
	Env
	GetAdditionalEnvs() []Env
}

func processEnvs(processedEnvs *[]Env, env TopLevelEnv, path string) {
	benchmarker := miniBenchmarker.GetInstance()
	log.Infof("Trying %s", env.Name())
	benchmarker.StartStage("env.ProcessPath.TryRespond")

	extraPaths, err := env.TryRespond(path)
	benchmarker.EndStage("env.ProcessPath.TryRespond")

	for _, extraPath := range extraPaths {
		log.Infof("Launching extraEnv")
		processEnvs(processedEnvs, extraPath.Env, extraPath.Path)
	}

	if err != nil {
		log.Warnf("tryrespond failed: %s", err.Error())
		return
	}

	log.Infof("OK response")

	*processedEnvs = append(*processedEnvs, env)

	for _, additionalEnv := range env.GetAdditionalEnvs() {
		benchmarker.StartStage("env.ProcessPath.TryRespondAdditionalEnv")
		log.Infof("trying additional env %s", additionalEnv.Name())
		_, err := additionalEnv.TryRespond(path, env)
		benchmarker.EndStage("env.ProcessPath.TryRespondAdditionalEnv")
		if err != nil {
			log.Infof("Fail response: %s", err.Error())
			continue
		}
		*processedEnvs = append(*processedEnvs, additionalEnv)
	}
}

func ProcessPath(path string, theDevFile data.DevfileData) error {
	benchmarker := miniBenchmarker.GetInstance()
	benchmarker.StartStage("env.ProcessPath")
	theEnvs := GetEnvs()
	processedEnvs := []Env{}
	//didRespondAtLeastOnce := false
	//hasMultipleEnvs := false

	for _, env := range theEnvs {
		processEnvs(&processedEnvs, env, path)
	}

	for _, env := range processedEnvs {
		benchmarker.StartStage("env.ProcessPath.Build")
		log.Infof("Building %s", env.Name())

		err := env.Build(theDevFile)
		benchmarker.EndStage("env.ProcessPath.Build")
		if err != nil {
			return err
		}
	}

	/*if hasMultipleEnvs || !didRespondAtLeastOnce { // todo: in this case maybe try go one level deeper as Ernst said
		log.Infof("Using default image (kitchen sink)")
		defaultEnv := GetDefaultEnv()
		_, err := defaultEnv.TryRespond(path)
		if err != nil {
			return errors.New("failed to apply kitchen sink")
		}
		err = defaultEnv.Build(theDevFile)
		if err != nil {
			return err
		}
	} else {
		for _, env := range processedEnvs {
			benchmarker.StartStage("env.ProcessPath.Build")
			log.Infof("Building %s", env.Name())

			err := env.Build(theDevFile)
			benchmarker.EndStage("env.ProcessPath.Build")
			if err != nil {
				return err
			}
		}
	}*/

	return nil
}
