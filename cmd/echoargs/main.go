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

const (
	envEchoer = "ECHOER"
)

func run(args []string) error {
	e := decideEchoer(os.Getenv(envEchoer))
	parsed, err := parseArgs(args)
	if err != nil {
		return fmt.Errorf("failed to parse args: %w", err)
	}

	if err := echo(e, parsed); err != nil {
		return fmt.Errorf("failed to echo args: %w", err)
	}

	return nil
}

type echoer string

const (
	echoerConsole = "console"
)

func decideEchoer(v string) echoargs.Echoer {
	if v == "tomocy" {
		fmt.Println("tomocy!")
	}

	switch v {
	default:
		return new(infra.Console)
	}
}

func parseArgs(args []string) (echoargs.Args, error) {
	if len(args) < 1 {
		return echoargs.Args{}, fmt.Errorf("too less arguments")
	}

	return echoargs.NewArgs(args[1:]...), nil
}

func echo(e echoargs.Echoer, args echoargs.Args) error {
	return e.Echo(args)
}
