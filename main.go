package main

import (
	"fmt"
	"os"
	"os/signal"
	"taskmaster/job_control"
	"taskmaster/parse_config"
)

func main() {
	if len(os.Args) < 2 {
		panic("1 argument expected")
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel)

	jobsConf, err := parse_config.ParseConfig(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(jobsConf)

	jC := job_control.NewJobControl()
	go jC.HandleSignals(signalChannel)

	//time.Sleep(20 * time.Second)
}
