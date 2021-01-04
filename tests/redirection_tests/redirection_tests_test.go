package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"taskmaster/task_master"
)

func captureOutput(f func()) (stdOut, stdErr string) {
	rOut, wOut, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	rErr, wErr, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = wOut
	os.Stderr = wErr
	out := make(chan string)
	outErr := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, rOut)
		out <- buf.String()
	}()

	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, rErr)
		outErr <- buf.String()
	}()
	wg.Wait()
	f()
	wOut.Close()
	wErr.Close()
	return <-out, <-outErr
}

func TestRedirection(t *testing.T) {
	stdOutRes, stdErrRes := captureOutput(func() {
		tM, err := task_master.NewTaskMaster("./redirection_tests.yaml", nil)
		if err != nil {
			t.Errorf("Redirection failed: %v\n", err)
		}
		if err = tM.Start(); err != nil {
			t.Errorf("Redirection failed: %v\n", err)
		}
		time.Sleep(25 * time.Millisecond)
		if err = tM.Stop(); err != nil {
			t.Errorf("Redirection failed: %v\n", err)
		}
	})

	// Test for regular redirection
	if stdErrRes != "stderr output\n" {
		t.Errorf("Redirection file failed: stderr output is not correct\n")
	}
	if stdOutRes != "stdout output\n" {
		t.Errorf("Redirection file failed: stderr output is not correct\n")
	}

	//Test for file redirection
	data, err := ioutil.ReadFile("./redirection_tests.stderr")
	if err != nil {
		t.Errorf("Redirection failed: %v\n", err)
	}
	if string(data) != "stderr output\n" {
		t.Errorf("Redirection file failed: stderr output is not correct\n")
	}
	data, err = ioutil.ReadFile("./redirection_tests.stdout")
	if err != nil {
		t.Errorf("Redirection failed: %v\n", err)
	}
	if string(data) != "stdout output\n" {
		t.Errorf("Redirection file failed: stdout output is not correct\n")
	}
}
