package task

import (
	"fmt"
	"log"
)

func (t *Task) GetStatus() string {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.status
}

func (t *Task) GetPid() (int, error) {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.cmd.Process == nil {
		return 0, fmt.Errorf("process is not started yet")
	}
	return t.cmd.Process.Pid, nil
}

func (t *Task) PrintStatus() {
	t.mux.Lock()
	defer t.mux.Unlock()
	pState := t.cmd.ProcessState.String()
	if t.cmd.ProcessState == nil {
		pState = "running"
	}
	log.Println(t.cmd.Process.Pid, pState)
}

func (t *Task) IsRunning() bool {
	t.mux.Lock()
	defer t.mux.Unlock()
	return t.isRunning
}
