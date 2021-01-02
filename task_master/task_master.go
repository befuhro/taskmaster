package task_master

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"taskmaster/task_master/tasks"
)

type TaskMaster struct {
	confPath string
	tasks    tasks.Tasks
}

func NewJobControl(confPath string) *TaskMaster {
	return &TaskMaster{
		confPath: confPath,
	}
}

func (j *TaskMaster) LoadConfig() error {
	data, err := ioutil.ReadFile(j.confPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &j.tasks)
}

func (j *TaskMaster) HandleSignals(signalChannel chan os.Signal) {
	for {
		sig := <-signalChannel
		for _, task := range j.tasks.Tasks {
			if sig.String() == task.StopSignal {
				fmt.Println(task)
				fmt.Println()
			}
		}
	}
}