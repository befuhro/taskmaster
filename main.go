package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"taskmaster/task_master"
)

func main() {
	if len(os.Args) < 2 {
		panic("1 argument expected")
	}

	fmt.Printf("%v\n\n", os.Getpid())

	signalChannel := make(chan os.Signal, 10)
	signal.Notify(signalChannel)

	tM, err := task_master.NewTaskMaster(os.Args[1])
	if err != nil {
		panic(err)
	}

	go tM.HandleSignals(signalChannel)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		cmd, _ := reader.ReadString('\n')
		if stop := tM.HandleCmd(cmd); stop {
			break
		}
	}
}
