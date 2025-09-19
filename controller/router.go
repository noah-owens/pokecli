// Package controller parses CLI arguments and dispatches them to the appropriate command handlers.
package controller

import (
	"fmt"
	"os"
	"pokecli/commands"
	"pokecli/format"
)

// Route dispatches CLI arguments to the appropriate command based on top-level and subcommand structure.
func Route(args []string) {

	// dispatch top-level command
	switch args[1] {
	case "get":
		// handle 'get' subcommands
		switch args[2] {
		case "pokemon":
			commands.GetPokemon(os.Args[3:])
		// unknown 'get' subcommand
		default:
			argsOutput := ""
			for _, a := range args[2:] {
				argsOutput += a + " "
			}
			fmt.Println(format.Red + "Unknown get subcommand: " + argsOutput)
			fmt.Println("Usage: pokecli get [pokemon/berry/ability] [arg]" + format.Reset)
			os.Exit(1)
		}
	// unknown top-level command
	default:
		argsOutput := ""
		for _, a := range args[1:] {
			argsOutput += a + " "
		}
		fmt.Println(format.Red + "Unknown command: " + argsOutput + format.Reset)
		os.Exit(1)
	}
}
