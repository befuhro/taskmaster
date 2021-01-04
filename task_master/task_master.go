package task_master

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"taskmaster/task_master/task"
)

type TaskMaster struct {
	wg       *sync.WaitGroup
	stopChan chan bool
	confPath string
	mux      sync.Mutex
	Tasks    map[string]*task.Task `yaml:"programs"`
}

func NewTaskMaster(confPath string, wg *sync.WaitGroup) (t *TaskMaster, err error) {
	t = &TaskMaster{confPath: confPath, stopChan: make(chan bool, 2), wg: wg}
	err = t.loadConfig()
	return
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
				t.HandleSIG(sig.String())
			}
		case <-t.stopChan:
			break
		}
	}
	t.wg.Done()
	log.Println("exit HandleSignals")
}