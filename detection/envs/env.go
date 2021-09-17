package envs

import (
	"errors"
	"github.com/devfile/devrunner/detection/miniBenchmarker"
	"github.com/devfile/library/pkg/devfile/parser/data"
	log "github.com/sirupsen/logrus"
)

func GetEnvs() []TopLevelEnv {
	return []TopLevelEnv{
		&javaMavenEnv{},
		&goEnv{},
	}
}

func GetDefaultEnv() TopLevelEnv {
	return &kitchenSinkImageEnv{}
}

type Env interface {
	Name() string
	TryRespond(rootPath string, additionalParams ...interface{}) error
	Build(devfile data.DevfileData) error
}

type TopLevelEnv interface {
	Env
	GetAdditionalEnvs() []Env
}

// ProcessPath is tested at realProjectTests

func ProcessPath(path string, theDevFile data.DevfileData) error {
	benchmarker := miniBenchmarker.GetInstance()
	benchmarker.StartStage("env.ProcessPath")
	theEnvs := GetEnvs()
	var processedEnvs []Env
	didRespondAtLeastOnce := false
	hasMultipleEnvs := false

	for _, env := range theEnvs {
		log.Infof("Trying %s", env.Name())
		benchmarker.StartStage("env.ProcessPath.TryRespond")

		err := env.TryRespond(path, nil)
		benchmarker.EndStage("env.ProcessPath.TryRespond")
		if err != nil {
			log.Infof("Fail response: %s", err.Error())
			continue
		}

		log.Infof("OK response")

		if didRespondAtLeastOnce && !hasMultipleEnvs {
			log.Infof("Will use the kitchen sink image because we found several envs")
			hasMultipleEnvs = true
			break
		}

		processedEnvs = append(processedEnvs, env)

		for _, additionalEnv := range env.GetAdditionalEnvs() {
			benchmarker.StartStage("env.ProcessPath.TryRespondAdditionalEnv")
			log.Infof("trying additional env %s", additionalEnv.Name())
			err := additionalEnv.TryRespond(path, env)
			benchmarker.EndStage("env.ProcessPath.TryRespondAdditionalEnv")
			if err != nil {
				log.Infof("Fail response: %s", err.Error())
				continue
			}
			processedEnvs = append(processedEnvs, additionalEnv)
		}

		didRespondAtLeastOnce = true
	}

	if hasMultipleEnvs || !didRespondAtLeastOnce { // todo: in this case maybe try go one level deeper as Ernst said
		log.Infof("Using default image (kitchen sink)")
		defaultEnv := GetDefaultEnv()
		err := defaultEnv.TryRespond(path)
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
	}

	return nil
}
