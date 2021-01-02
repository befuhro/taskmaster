package task_master

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (t *TaskMaster) loadConfig() error {
	data, err := ioutil.ReadFile(t.confPath)
	if err != nil {
		return err
	}
	if err =  yaml.Unmarshal(data, &t.tasks); err != nil {
		return err
	}
	return nil
}

func (t *TaskMaster) reloadConfig() error {
	fmt.Println("reload config")
	return nil
}