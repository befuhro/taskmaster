package task

import (
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Task struct {
	mux          sync.Mutex
	cmd          *exec.Cmd
	isRunning    bool
	status       string
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
}

func (t *Task) Stop() error {
	if err := t.terminate(); err != nil {
		return err
	}
	time.AfterFunc(time.Duration(t.StopTime)*time.Second, t.kill)
	return nil
}

func (t *Task) Start() error {
	t.mux.Lock()
	args := strings.Split(t.Cmd, " ")
	t.cmd = exec.Command(args[0], args[1:]...)
	t.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	t.mux.Unlock()
	if err := t.redirectOutput(); err != nil {
		return err
	}
	// Run proc after given start time * second if autostart is on
	if t.AutoStart {
		time.AfterFunc(time.Duration(t.StartTime)*time.Second, t.asyncRun)
	}
	return nil
}
