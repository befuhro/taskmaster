package main

import (
	"fmt"
	"os"
	"sync"

	"taskmaster/task_master"
)

func main() {
	if len(os.Args) < 2 {
		panic("1 argument expected")
	}
	wg := sync.WaitGroup{}
	tM, err := task_master.NewTaskMaster(os.Args[1], &wg)
	if err != nil {
		panic(err)
	}

	if err = tM.Start(); err != nil {
		panic(err)
	}

	wg.Add(2)
	go tM.WaitSignals()
	go tM.WaitCmd()

	wg.Wait()
	fmt.Println("taskmaster is closed")
}
