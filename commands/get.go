// Package commands defines CLI command handlers for fetching and displaying Pokémon data.
package commands

import (
	"fmt"
	"os"
	"slices"

	"pokecli/api"
	"pokecli/format"
)

// GetPokemon handles the 'get pokemon' CLI command.
// Fetches data from PokeAPI and renders either a brief or detailed summary.
// Flags:
//
//	--detailed / -d : show stats, held items, and moves
//	--moves    / -m : show full moveset (only applies to --detailed)
func GetPokemon(args []string) {

	// Missing required Pokémon name
	if len(args) < 1 {
		fmt.Println(format.Red + "Usage: pokecli get pokemon [name] [--detailed] [--moves]" + format.Reset)
		os.Exit(1)
	}

	name := args[0]
	detailed := false
	allMoves := false

	// Supported flags
	detailedFlags := []string{"--detailed", "-d"}
	movesFlags := []string{"--moves", "-m"}

	// Parse flags
	for _, arg := range args[1:] {
		if slices.Contains(detailedFlags, arg) {
			detailed = true
		}

		if slices.Contains(movesFlags, arg) {
			allMoves = true
		}
	}

	// Fetch Pokémon data
	data, err := api.FetchPokemon(name)
	if err != nil {
		fmt.Printf(format.Red+"Error fetching data: %s"+format.Reset, err)
		os.Exit(1)
	}

	// Transform and render summary
	summary := format.PokemonToSummary(data)

	if detailed {
		format.PrintDetailedPokemonSummary(summary, allMoves)
	} else {
		format.PrintPokemonSummary(summary)

		// Warn if --moves used without --detailed
		if allMoves {
			fmt.Println(format.Yellow + "Warning: [--moves] only applies to detailed summary")
			fmt.Println("Usage: pokecli get pokemon --detailed --moves" + format.Reset)
		}
	}
}
