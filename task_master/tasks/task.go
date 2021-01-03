package tasks

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
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
	outFile, err := os.Create(t.StdOut)
	if err != nil {
		return err
	}
	errFile, err := os.Create(t.StdErr)
	if err != nil {
		return err
	}
	cmd.Stderr = errFile
	cmd.Stdout = outFile
	return nil
}

func (t *Task) PrintStatus() {
	pState := t.cmd.ProcessState.String()
	if t.cmd.ProcessState == nil {
		pState = "running"
	}
	log.Println(t.cmd.Process.Pid, pState)
}

func (t *Task) Stop() error {
	return syscall.Kill(-t.cmd.Process.Pid, syscall.SIGKILL)
}

func (t *Task) Start() error {
	args := strings.Split(t.Cmd, " ")
	t.cmd = exec.Command(args[0], args[1:]...)
	t.cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := t.redirectOutput(t.cmd); err != nil {
		return err
	}
	log.Println(args)
	go t.cmd.Run()
	return nil
}
