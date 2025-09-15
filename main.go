package main

import (
	"fmt"
	"os"

	"pokecli/commands"
	"pokecli/format"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(format.Red + "Usage: pokecli [command] [args]" + format.Reset)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		commands.Get(os.Args[2:])
	default:
		fmt.Println(format.Red + "Unknown command: " + os.Args[1] + format.Reset)
		os.Exit(1)
	}
}
