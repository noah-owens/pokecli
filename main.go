// Package main is the entry point to pokecli.
// It parses CLI arguments and delegates routing to the controller package.
package main

import (
	"fmt"
	"os"

	"pokecli/controller"
	"pokecli/format"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(format.Red + "Usage: pokecli [command] [args]" + format.Reset)
		os.Exit(1)
	}

	controller.Route(os.Args)
}
