package task_master

import (
	"fmt"
	"os"
	"syscall"

	"taskmaster/task_master/tasks"
)

type TaskMaster struct {
	confPath string
	tasks    tasks.Tasks
}

func NewTaskMaster(confPath string) (t *TaskMaster, err error) {
	t = &TaskMaster{confPath: confPath}
	err = t.loadConfig()
	if err != nil {
		return
	}
	err = t.tasks.Start()
	return
}

func (t *TaskMaster) HandleSignals(signalChannel chan os.Signal) {
	for {
		sig := <-signalChannel
		if sig == syscall.SIGHUP {
			t.reloadConfig()
		} else {
			t.tasks.HandleSIG(sig.String())
		}
	}
}

func (t *TaskMaster) stop() {
	fmt.Println("STOP")
	t.tasks.Stop()
}