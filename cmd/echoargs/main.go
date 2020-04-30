package main

import (
	"fmt"
	"os"

	"github.com/tomocy/go-echoargs"
	"github.com/tomocy/go-echoargs/infra"
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

func decideEchoer(v string) echoargs.Echoer {
	switch v {
	default:
		return new(infra.Console)
	}
}
