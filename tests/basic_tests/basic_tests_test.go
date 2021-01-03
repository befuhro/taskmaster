package main

import (
	"testing"
	"time"

	"taskmaster/task_master"
	"taskmaster/task_master/tasks"
)

func TestConfigurationLoading(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./basic_tests.yaml", nil)
	if err != nil {
		t.Errorf("Configuration Loading failed: %v\n", err)
	}
	task, ok := tM.GetTask("basic_tests")
	if !ok {
		t.Errorf("Configuration Loading failed: task 'basic_tests' was not created")
	}
	expectedTask := tasks.Task{
		Cmd:          "./basic_tests",
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
		StdOut:       "./basic_tests.stdout",
		StdErr:       "./basic_tests.stderr",
		Env:          map[string]string{"STARTED_BY": "taskmaster", "ANSWER":"42"},
	}
	if !tasks.TaskCmp(&expectedTask, task) {
		t.Errorf("Configuration Loading failed: task 'basic_tests' was not loaded properly")
	}
}

func TestStartAndStopTask(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./basic_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(time.Second)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
}
