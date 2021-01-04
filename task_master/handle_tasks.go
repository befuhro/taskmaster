package task_master

import (
	"fmt"
	"log"

	"taskmaster/task_master/task"
)

func (t *TaskMaster) Start() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, tsk := range t.Tasks {
		if err := tsk.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (t *TaskMaster) Stop() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, tsk := range t.Tasks {
		if err := tsk.Stop(); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (t *TaskMaster) StopTask(taskName string) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	tsk, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return tsk.Stop()
}

func (t *TaskMaster) StartTask(taskName string) error {
	t.mux.Lock()
	defer t.mux.Unlock()
	tsk, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return tsk.Start()
}

func (t *TaskMaster) HandleSIG(sig string) {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, tsk := range t.Tasks {
		if sig == tsk.StopSignal {
			if err := tsk.Stop(); err != nil {
				log.Println(err)
			}
		}
	}
}

func (t *TaskMaster) PrintStatus() {
	t.mux.Lock()
	defer t.mux.Unlock()
	for _, tsk := range t.Tasks {
		tsk.PrintStatus()
	}
}

func (t *TaskMaster) GetTaskStatus(taskName string) (string, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	tsk, ok := t.Tasks[taskName]
	if !ok {
		return "", fmt.Errorf("task '%v' does not exist", taskName)
	}
	return tsk.GetStatus(), nil
}

func (t *TaskMaster) IsTaskRunning(taskName string) (bool, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	tsk, ok := t.Tasks[taskName]
	if !ok {
		return false, fmt.Errorf("task '%v' does not exist", taskName)
	}
	return tsk.IsRunning(), nil
}

func (t *TaskMaster) GetTaskPid(taskName string) (int, error) {
	tsk, ok := t.Tasks[taskName]
	if !ok {
		return 0, fmt.Errorf("task '%v' does not exist", taskName)
	}
	return tsk.GetPid()
}

func (t *TaskMaster) GetTask(taskName string) (*task.Task, bool) {
	tsk, ok := t.Tasks[taskName]
	return tsk, ok
}