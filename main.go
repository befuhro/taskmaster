package main

import (
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

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel)

	jC := task_master.NewJobControl(os.Args[1])
	if err := jC.LoadConfig(); err != nil {
		panic(err)
	}
	jC.HandleSignals(signalChannel)
}
