package tasks

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
}

