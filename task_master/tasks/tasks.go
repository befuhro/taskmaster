package tasks

import (
	"fmt"
	"log"
	"sync"

	"taskmaster/task_master/tasks/task"
)

type Tasks struct {
	mux   sync.Mutex
	Tasks map[string]*task.Task `yaml:"programs"`
}

func (t *Tasks) Start() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, task := range t.Tasks {
		if err := task.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (t *Tasks) Stop() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, task := range t.Tasks {
		if err := task.Stop(); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (t *Tasks) StopTask(taskName string) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	task, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.Stop()
}

func (t *Tasks) StartTask(taskName string) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	task, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.Start()
}

func (t *Tasks) HandleSIG(sig string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, task := range t.Tasks {
		if sig == task.StopSignal {
			if err := task.Stop(); err != nil {
				log.Println(err)
			}
		}
	}
}

func (t *Tasks) PrintStatus() {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, task := range t.Tasks {
		task.PrintStatus()
	}
}

func (t *Tasks) GetTaskStatus(taskName string) (string, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	task, ok := t.Tasks[taskName]
	if !ok {
		return "", fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.GetStatus(), nil
}

func (t *Tasks) IsTaskRunning(taskName string) (bool, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	task, ok := t.Tasks[taskName]
	if !ok {
		return false, fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.IsRunning(), nil
}