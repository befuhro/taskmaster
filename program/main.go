package main

import (
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for {
		log.Println("\t", os.Getpid(), "started at:", start.Format("15:04:05.00000"))
		time.Sleep(1 * time.Second)
	}
}