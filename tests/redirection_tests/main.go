package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stderr, "stderr output\n")
	fmt.Fprint(os.Stdout, "stdout output\n")
}