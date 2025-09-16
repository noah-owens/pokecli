package controller

import (
	"fmt"
	"os"
	"pokecli/commands"
	"pokecli/format"
)

func Route(args []string) {
	switch args[1] {
	case "get":
		switch args[2] {
		case "pokemon":
			commands.GetPokemon(os.Args[3:])
		default:
			argsOutput := ""
			for _, a := range args[2:] {
				argsOutput += a + " "
			}
			fmt.Println(format.Red + "Unknown get subcommand: " + argsOutput)
			fmt.Println("Usage: pokecli get [pokemon/berry/ability] [arg]" + format.Reset)
			os.Exit(1)
		}
	default:
		argsOutput := ""
		for _, a := range args[1:] {
			argsOutput += a + " "
		}
		fmt.Println(format.Red + "Unknown command: " + argsOutput + format.Reset)
		os.Exit(1)
	}
}
