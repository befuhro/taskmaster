package task_master

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"taskmaster/task_master/tasks"
)

type TaskMaster struct {
	wg       *sync.WaitGroup
	stopChan chan bool
	confPath string
	tasks    tasks.Tasks
}

func NewTaskMaster(confPath string, wg *sync.WaitGroup) (t *TaskMaster, err error) {
	t = &TaskMaster{confPath: confPath, stopChan: make(chan bool, 2), wg: wg}
	err = t.loadConfig()
	return
}

func (t *TaskMaster) StartTasks() error {
	return t.tasks.Start()
}

func (t *TaskMaster) GetTask(taskName string) (*tasks.Task, bool) {
	task, ok := t.tasks.Tasks[taskName]
	return task, ok
}

func (t *TaskMaster) WatchTask() {
	for {
		select {
		case <-t.stopChan:
			break
		default:
			log.Println("watch tasks")
			time.Sleep(time.Second)
		}
	}
	log.Println("exit WatchTask")
}

func (t *TaskMaster) WaitSignals() {
	signalChannel := make(chan os.Signal, 10)
	signal.Notify(signalChannel)

	for {
		select {
		case sig := <-signalChannel:
			if sig == syscall.SIGHUP {
				t.reloadConfig()
			} else {
				t.tasks.HandleSIG(sig.String())
			}
		case <-t.stopChan:
			break
		}
	}
	t.wg.Done()
	log.Println("exit HandleSignals")
}
