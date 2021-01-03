package tasks

import (
	"fmt"
	"log"
)

type Tasks struct {
	Tasks map[string]*Task `yaml:"programs"`
}

func (t *Tasks) Start() error {
	for _, task := range t.Tasks {
		if err := task.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (t *Tasks) Stop() error {
	for _, task := range t.Tasks {
		if err := task.Stop(); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (t *Tasks) StopTask(taskName string) error {
	task, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.Stop()
}

func (t *Tasks) StartTask(taskName string) error {
	task, ok := t.Tasks[taskName]
	if !ok {
		return fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.Start()
}

func (t *Tasks) HandleSIG(sig string) {
	for _, task := range t.Tasks {
		if sig == task.StopSignal {
			if err := task.Stop(); err != nil {
				log.Println(err)
			}
		}
	}
}

func (t *Tasks) PrintStatus() {
	for _, task := range t.Tasks {
		task.PrintStatus()
	}
}

func (t *Tasks) GetTaskStatus(taskName string) (string, error) {
	task, ok := t.Tasks[taskName]
	if !ok {
		return "", fmt.Errorf("task '%v' does not exist", taskName)
	}
	return task.GetStatus(), nil
}