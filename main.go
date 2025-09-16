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
		switch os.Args[2] {
		case "pokemon":
			commands.GetPokemon(os.Args[3:])
		default:
			argsOutput := ""
			for _, arg := range os.Args[2:] {
				argsOutput += arg + " "
			}
			fmt.Println(format.Red + "Unknown get argument: " + argsOutput + format.Reset)
			os.Exit(1)
		}
	default:
		argsOutput := ""
		for _, arg := range os.Args[1:] {
			argsOutput += arg + " "
		}
		fmt.Println(format.Red + "Unknown command: " + argsOutput + format.Reset)
		os.Exit(1)
	}
}
