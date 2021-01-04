package task_master

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"taskmaster/task_master/task"
)

func (t *TaskMaster) loadConfig() error {
	data, err := ioutil.ReadFile(t.confPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &t); err != nil {
		return err
	}
	return nil
}

func (t *TaskMaster) reloadConfig() error {
	data, err := ioutil.ReadFile(t.confPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &t); err != nil {
		return err
	}
	for taskName, tsk := range t.Tasks {
		oldTask, ok := t.Tasks[taskName]
		if !ok || !task.TaskCmp(oldTask, tsk) {
			if err = oldTask.Stop(); err != nil {
				log.Println(err)
			}
			if err = tsk.Start(); err != nil {
				log.Println(err)
			}
			t.Tasks[taskName] = tsk
		}
	}
	return nil
}