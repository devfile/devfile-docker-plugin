package miniBenchmarker

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var benchmarkerInstance *miniBenchmarker

func GetInstance() *miniBenchmarker {
	if benchmarkerInstance == nil {
		if benchmarkerInstance == nil {
			benchmarkerInstance = &miniBenchmarker{
				stages: map[string]int64{},
			}
		}
	}

	return benchmarkerInstance
}

type miniBenchmarker struct {
	stages map[string]int64
}

func (*miniBenchmarker) logInfo(entry string) {
	log.Infof("[miniBenchmarker] %s", entry)
}

func (*miniBenchmarker) logError(entry string) {
	log.Errorf("[miniBenchmarker] %s", entry)
}

func (this *miniBenchmarker) StartStage(stageName string) {
	this.logInfo("starting " + stageName)
	value, exists := this.stages[stageName]
	if exists && value != -1 {
		this.logError("this stage has been started")
		return
	}

	this.stages[stageName] = time.Now().UnixNano()
}

func (this *miniBenchmarker) EndStage(stageName string) {
	_, exists := this.stages[stageName]
	if !exists {
		this.logError("this stage has NOT been started")
		return
	}

	this.stages[stageName] = -1

	takenTime := (time.Now().UnixNano() - this.stages[stageName]) / 1e+6

	this.logInfo(stageName + " took " + strconv.FormatInt(takenTime, 10) + "ns")
}
