package main

import (
	"taskmaster/task_master/tasks"
	"testing"

	"taskmaster/task_master"
)

func TestConfigurationLoading(t *testing.T) {
	tM, err := task_master.NewTaskMaster("config/test.yaml", nil)
	if err != nil {
		t.Errorf("Configuration Loading failed: %v\n", err)
	}
	task, ok := tM.GetTask("test")
	if !ok {
		t.Errorf("Configuration Loading failed: task 'test' was not created")
	}
	expectedTask := tasks.Task{
		Cmd:          "./program/program",
		NumProcs:     1,
		UMask:        18,
		WorkingDir:   "/tmp",
		AutoStart:    true,
		AutoRestart:  "unexpected",
		ExitCodes:    []int{0,2},
		StartRetries: 3,
		StartTime:    0,
		StopSignal:   "terminated",
		StopTime:     10,
		StdOut:       "./output/test.stdout",
		StdErr:       "./output/test.stderr",
		Env:          map[string]string{"STARTED_BY": "taskmaster", "ANSWER":"42"},
	}
	if !tasks.TaskCmp(&expectedTask, task) {
		t.Errorf("Configuration Loading failed: task 'test' was not loaded properly")

	}
}

//func TestStartTasks(t *testing.T) {
//	tM, err := task_master.NewTaskMaster("./config/config.yaml", nil)
//	if err != nil {
//		t.Errorf("Start Tasks failed: %v\n", err)
//	}
//	if err = tM.StartTasks(); err != nil {
//		t.Errorf("Start Tasks failed: %v\n", err)
//	}
//	time.Sleep(time.Second)
//}
