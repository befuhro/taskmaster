package task_master

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var regStop, _ = regexp.Compile("^stop")
var regStart, _ = regexp.Compile("^start")

func (t *TaskMaster) WaitCmd() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		cmd, _ := reader.ReadString('\n')
		if cmd == "status\n" {
			t.tasks.PrintStatus()
		} else if cmd == "stop\n" {
			t.tasks.Stop()
			t.wg.Done()
			close(t.stopChan)
			break
		} else if regStop.MatchString(cmd) {
			t.stopCmd(cmd)
		} else if regStart.MatchString(cmd) {
			t.startCmd(cmd)
		}
	}
	t.wg.Done()
}

func (t *TaskMaster) handleCmd(cmd string) (stop bool) {
	if cmd == "status\n" {
		t.tasks.PrintStatus()
	} else if regStop.MatchString(cmd) {
		t.stopCmd(cmd)
	} else if regStart.MatchString(cmd) {
		t.startCmd(cmd)
	}
	return
}

func (t *TaskMaster) stopCmd(cmd string) {
	args := strings.Split(cmd, " ")
	if len(args) < 2 {
		t.tasks.Stop()
	} else {
		taskName := args[1][:len(args[1])-1]
		if err := t.tasks.StopTask(taskName); err != nil {
			log.Println("could not stop task", err)
		} else {
			log.Println("stopped task", taskName)
		}
	}
}

func (t *TaskMaster) startCmd(cmd string) {
	args := strings.Split(cmd, " ")
	if len(args) >= 2 {
		taskName := args[1][:len(args[1])-1]
		if err := t.tasks.StartTask(taskName); err != nil {
			log.Println("could not start task", err)
		} else {
			log.Println("started task", taskName)
		}
	}
}
