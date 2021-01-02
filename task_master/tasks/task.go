package tasks

import (
	"os"
	"os/exec"
	"strings"
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
	outFile		*os.File
	errFile		*os.File

}

func (t *Task) redirectOutput() error {
	outFile, err := os.Create(t.StdOut)
	if err != nil {
		return err
	}
	errFile, err := os.Create(t.StdErr)
	if err != nil {
		return err
	}
	t.outFile = outFile
	t.errFile = errFile
	return nil
}

func (t *Task) Start() error {
	if err := t.redirectOutput(); err != nil {
		return err
	}
	args := strings.Split(t.Cmd, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = t.errFile
	cmd.Stdout = t.outFile
	if err := cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}