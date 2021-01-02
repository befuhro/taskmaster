package tasks

import (
	"fmt"
	"log"
)

type Tasks struct {
	Tasks map[string]Task `yaml:"programs"`
}

func (t *Tasks) Start() error {
	for _, task := range t.Tasks {
		if err := task.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (t *Tasks) HandleSIG(sig string) {
	for _, task := range t.Tasks {
		if sig == task.StopSignal {
			fmt.Println(task)
			fmt.Println()
		}
	}
}

func (t *Tasks) PrintStatus() {
	log.Println("PrintStatus")
}
