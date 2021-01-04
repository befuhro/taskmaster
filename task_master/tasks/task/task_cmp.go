package task

import "reflect"

func TaskCmp(task1, task2 *Task) bool {
	if task1.Cmd != task2.Cmd ||
		task1.NumProcs != task2.NumProcs ||
		task1.UMask != task2.UMask ||
		task1.WorkingDir != task2.WorkingDir ||
		task1.AutoStart != task2.AutoStart ||
		task1.AutoRestart != task2.AutoRestart ||
		!reflect.DeepEqual(task1.ExitCodes, task2.ExitCodes) ||
		task1.StartRetries != task2.StartRetries ||
		task1.StartTime != task2.StartTime ||
		task1.StopSignal != task2.StopSignal ||
		task1.StopTime != task2.StopTime ||
		task1.StdOut != task2.StdOut ||
		task1.StdErr != task2.StdErr ||
		!reflect.DeepEqual(task1.Env, task2.Env) {
		return false
	}
	return true
}

