package task_master

func (t *TaskMaster) HandleCmd(cmd string) (stop bool) {
	if cmd == "status\n" {
		t.tasks.PrintStatus()
	} else if cmd == "stop\n" {
		t.stop()
		stop = true
	}
	return
}