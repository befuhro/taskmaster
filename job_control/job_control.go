package job_control

import (
	"log"
	"os"
)

type JobControl struct {

}

func NewJobControl() *JobControl {
	return &JobControl{}
}

func (j *JobControl) HandleSignals(signalChannel chan os.Signal) {
	for {
		sig := <-signalChannel
		log.Println(sig)
	}
}