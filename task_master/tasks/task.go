package tasks

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

type Task struct {
	Cmd          string            `yaml:"cmd"`
	NumProcs     uint8             `yaml:"numprocs"`
	UMask        uint8             `yaml:"umask"`
	WorkingDir   string            `yaml:"workingdir"`
	AutoStart    bool              `yaml:"autostart"`
	AutoRestart  string            `yaml:"autorestart"`
	ExitCodes    []int             `yaml:"exitcodes"`
	StartRetries uint              `yaml:"startretries"`
	StartTime    uint              `yaml:"starttime"`
	StopSignal   string            `yaml:"stopsignal"`
	StopTime     uint              `yaml:"stoptime"`
	StdOut       string            `yaml:"stdout"`
	StdErr       string            `yaml:"stderr"`
	Env          map[string]string `yaml:"env"`
	cmd          *exec.Cmd
}

func (t *Task) redirectOutput(cmd *exec.Cmd) error {
	if t.StdOut != "" {
		outFile, err := os.Create(t.StdOut)
		if err != nil {
			return err
		}
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = os.Stdout
	}
	if t.StdErr != "" {
		errFile, err := os.Create(t.StdErr)
		if err != nil {
			return err
		}
		cmd.Stderr = errFile
	} else {
		cmd.Stderr = os.Stderr
	}
	return nil
}

func (t *Task) GetStatus() string {
	pState := t.cmd.ProcessState.String()
	if t.cmd.ProcessState == nil {
		return "running"
	}
	return pState
}

func (t *Task) GetPid() (int, error) {
	if t.cmd.Process == nil {
		return 0, fmt.Errorf("process is not started yet")
	}
	return t.cmd.Process.Pid, nil
}

func (t *Task) PrintStatus() {
	pState := t.cmd.ProcessState.String()
	if t.cmd.ProcessState == nil {
		pState = "running"
	}
	log.Println(t.cmd.Process.Pid, pState)
}

func (t *Task) Stop() error {
	if t.cmd.Process == nil {
		return fmt.Errorf("process is not running")
	}
	// Kill process after stop time * second
	time.AfterFunc(time.Duration(t.StopTime) * time.Second, func() {
		if t.cmd.Process == nil {
			return
		}
		if err := syscall.Kill(-t.cmd.Process.Pid, syscall.SIGKILL); err != nil {
			log.Println(err)
		}
	})
	return syscall.Kill(-t.cmd.Process.Pid, syscall.SIGTERM)
}

func (t *Task) Start() error {
	args := strings.Split(t.Cmd, " ")
	t.cmd = exec.Command(args[0], args[1:]...)
	t.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := t.redirectOutput(t.cmd); err != nil {
		return err
	}

	// Run proc after given start time * second if autostart is on
	if t.AutoStart {
		go func(cmd *exec.Cmd) {
			time.AfterFunc(time.Duration(t.StartTime) * time.Second, func() {
				if err := t.cmd.Run(); err != nil {
					log.Println(err)
				}
			})
		}(t.cmd)
	}
	return nil
}
