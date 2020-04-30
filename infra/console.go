package infra

import (
	"fmt"

	"github.com/tomocy/go-echoargs"
)

type Console struct{}

func (c Console) Echo(args echoargs.Args) error {
	vs := args.Values()
	for i, v := range vs {
		fmt.Print(v)
		if i != len(vs)-1 {
			fmt.Print(", ")
		}
	}

	fmt.Println()

	return nil
}
