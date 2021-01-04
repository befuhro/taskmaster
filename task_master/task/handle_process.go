package task

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func (t *Task) redirectOutput() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.StdOut != "" {
		outFile, err := os.Create(t.StdOut)
		if err != nil {
			return err
		}
		t.cmd.Stdout = outFile
	} else {
		t.cmd.Stdout = os.Stdout
	}
	if t.StdErr != "" {
		errFile, err := os.Create(t.StdErr)
		if err != nil {
			return err
		}
		t.cmd.Stderr = errFile
	} else {
		t.cmd.Stderr = os.Stderr
	}
	return nil
}

func (t *Task) asyncRun() {
	t.mux.Lock()
	defer t.mux.Unlock()
	if err := t.cmd.Start(); err != nil {
		log.Println("asyncRun: ", err)
		return
	}
	t.isRunning = true
	t.status = "running"
	go t.wait()
}

// wait for end of process
func (t *Task) wait() {
	t.mux.Lock()
	cmd := *t.cmd
	t.mux.Unlock()
	err := cmd.Wait()
	t.mux.Lock()
	defer t.mux.Unlock()
	t.isRunning = false
	if err != nil {
		t.status = err.Error()
		log.Println(err)
		return
	}
	t.status = "stopped"
}

func (t *Task) terminate() error {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.cmd.Process == nil {
		return fmt.Errorf("process is not running")
	}
	if err := syscall.Kill(-t.cmd.Process.Pid, syscall.SIGTERM); err != nil {
		return err
	}
	return nil
}

func (t *Task) kill() {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.cmd.Process != nil {
		if err := syscall.Kill(-t.cmd.Process.Pid, syscall.SIGKILL); err != nil {
			log.Println(err)
		}
	}
}