package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	return fmt.Errorf("not implemented")
}

type echoer string

const (
	echoerConsole = "console"
)
