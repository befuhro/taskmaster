package task_master

import (
	"io/ioutil"
	"log"
	"reflect"

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
		if !ok || !tasksAreEqual(oldTask, task) {
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

func tasksAreEqual(oldTask, newTask *tasks.Task) bool {
	if oldTask.Cmd != newTask.Cmd ||
		oldTask.NumProcs != newTask.NumProcs ||
		oldTask.UMask != newTask.UMask ||
		oldTask.WorkingDir != newTask.WorkingDir ||
		oldTask.AutoStart != newTask.AutoStart ||
		oldTask.AutoRestart != newTask.AutoRestart ||
		!reflect.DeepEqual(oldTask.ExitCodes, newTask.ExitCodes) ||
		oldTask.StartRetries != newTask.StartRetries ||
		oldTask.StartTime != newTask.StartTime ||
		oldTask.StopSignal != newTask.StopSignal ||
		oldTask.StopTime != newTask.StopTime ||
		oldTask.StdOut != newTask.StdOut ||
		oldTask.StdErr != newTask.StdErr ||
		!reflect.DeepEqual(oldTask.Env, newTask.Env) {
		return false
	}
	return true
}