package main

import (
	"testing"
	"time"

	"taskmaster/task_master"
)

// Start process
func TestStart(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus != "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
}

// Start process by manually
func TestStartManually(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus != "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
}

// Stop process
func TestStop(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(200 * time.Millisecond)
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus == "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
}

// Stop process by manually
func TestStopManually(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(200 * time.Millisecond)
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus == "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
}


// Restart process
func TestRestart(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus != "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
}

// Restart process by manually
func TestRestartManually(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StopTasks(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.StartTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	} else if taskStatus != "running" {
		t.Errorf("Start Tasks failed: task status is '%v' and not 'running'\n", taskStatus)
	}
}