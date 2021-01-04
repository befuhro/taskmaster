package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"taskmaster/task_master"
)

func isTaskRunning(tM *task_master.TaskMaster) (bool, error) {
	taskStatus, err := tM.GetTaskStatus("start_and_stop_tests")
	if err != nil {
		return false, err
	} else if taskStatus != "running" {
		return false, fmt.Errorf("task status is '%v' and not 'running'\n", taskStatus)
	}
	taskPid, err := tM.GetTaskPid("start_and_stop_tests")
	if err != nil {
		return false, err
	}
	log.Println("PID:", taskPid, taskStatus)
	return true, nil
}

// Start process
func TestStart(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.Start(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}

	time.Sleep(25 * time.Millisecond)

	isRunning, err := isTaskRunning(tM)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if !isRunning {
		t.Errorf("Start Tasks failed: task is not running\n")
	}
	if err = tM.Stop(); err != nil {
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

	time.Sleep(25 * time.Millisecond)

	isRunning, err := isTaskRunning(tM)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if !isRunning {
		t.Errorf("Start Tasks failed: task is not running\n")
	}

	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
}

// Stop process
func TestStop(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.Start(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)

	isRunning, _ := isTaskRunning(tM)
	if isRunning {
		t.Errorf("Stop Tasks failed: task is runnning but be should terminated\n")
	}

	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
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
		t.Errorf("Stop Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)

	isRunning, _ := isTaskRunning(tM)
	if isRunning {
		t.Errorf("Stop Tasks failed: task is runnning but be should terminated\n")
	}
}

// Restart process
func TestRestart(t *testing.T) {
	tM, err := task_master.NewTaskMaster("./start_and_stop_tests.yaml", nil)
	if err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	if err = tM.Start(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.Start(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}

	time.Sleep(25 * time.Millisecond)
	isRunning, err := isTaskRunning(tM)
	if err != nil {
		t.Errorf("Restart Tasks failed: %v\n", err)
	}
	if !isRunning {
		t.Errorf("Start Tasks failed: task is not running\n")
	}

	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
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
	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	if err = tM.StartTask("start_and_stop_tests"); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
	time.Sleep(25 * time.Millisecond)
	isRunning, err := isTaskRunning(tM)
	if err != nil {
		t.Errorf("Restart Tasks failed: %v\n", err)
	}
	if !isRunning {
		t.Errorf("Start Tasks failed: task is not running\n")
	}

	if err = tM.Stop(); err != nil {
		t.Errorf("Start Tasks failed: %v\n", err)
	}
}
