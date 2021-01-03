package task_master

import (
	"log"
	"regexp"
	"strings"
)

var regStop, _ = regexp.Compile("^stop")
var regStart, _ = regexp.Compile("^start")

func (t *TaskMaster) HandleCmd(cmd string) (stop bool) {
	args := strings.Split(cmd, " ")
	if cmd == "status\n" {
		t.tasks.PrintStatus()
	} else if regStop.MatchString(cmd) {
		if len(args) < 2 {
			t.tasks.Stop()
			stop = true
		} else {
			taskName := args[1][:len(args[1])-1]
			if err := t.tasks.StopTask(taskName); err != nil {
				log.Println("could not stop task", err)
			} else {
				log.Println("stopped task", taskName)
			}
		}
	} else if regStart.MatchString(cmd) {
		if len(args) >= 2 {
			taskName := args[1][:len(args[1])-1]
			if err := t.tasks.StartTask(taskName); err != nil {
				log.Println("could not start task", err)
			} else {
				log.Println("started task", taskName)
			}
		}
	}
	return
}
