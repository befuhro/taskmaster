package task_master

import (
	"os"
	"os/signal"
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

func (t *TaskMaster) HandleSignals() {
	signalChannel := make(chan os.Signal, 10)
	signal.Notify(signalChannel)

	for {
		sig := <-signalChannel
		if sig == syscall.SIGHUP {
			t.reloadConfig()
		} else {
			t.tasks.HandleSIG(sig.String())
		}
	}
}