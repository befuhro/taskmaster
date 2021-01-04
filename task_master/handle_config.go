package task_master

import (
	"io/ioutil"
	"log"
	task2 "taskmaster/task_master/tasks/task"

	"gopkg.in/yaml.v2"

	"taskmaster/task_master/tasks"
)

func (t *TaskMaster) loadConfig() error {
	data, err := ioutil.ReadFile(t.confPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, &t.tasks); err != nil {
		return err
	}
	return nil
}

func (t *TaskMaster) reloadConfig() error {
	data, err := ioutil.ReadFile(t.confPath)
	if err != nil {
		return err
	}
	var loadedTasks tasks.Tasks
	if err = yaml.Unmarshal(data, &loadedTasks); err != nil {
		return err
	}
	for taskName, task := range loadedTasks.Tasks {
		oldTask, ok := t.tasks.Tasks[taskName]
		if !ok || !task2.TaskCmp(oldTask, task) {
			if err = oldTask.Stop(); err != nil {
				log.Println(err)
			}
			if err = task.Start(); err != nil {
				log.Println(err)
			}
			t.tasks.Tasks[taskName] = task
		}
	}
	return nil
}