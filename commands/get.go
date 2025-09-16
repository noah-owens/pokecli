package commands

import (
	"fmt"
	"os"
	"slices"

	"pokecli/api"
	"pokecli/format"
)

func GetPokemon(args []string) {
	if len(args) < 1 {
		fmt.Println(format.Red + "Usage: pokecli get pokemon [name] [--detailed] [--moves]" + format.Reset)
		os.Exit(1)
	}

	name := args[0]
	detailed := false
	allMoves := false

	detailedFlags := []string{"--detailed", "-d"}
	movesFlags := []string{"--moves", "-m"}

	for _, arg := range args[1:] {
		if slices.Contains(detailedFlags, arg) {
			detailed = true
		}

		if slices.Contains(movesFlags, arg) {
			allMoves = true
		}
	}

	data, err := api.FetchPokemon(name)
	if err != nil {
		fmt.Printf(format.Red+"Error fetching data: %s"+format.Reset, err)
		os.Exit(1)
	}

	summary := api.PokemonToSummary(data)

	if detailed {
		format.PrintDetailedPokemonSummary(summary, allMoves)
	} else {
		format.PrintPokemonSummary(summary)

		if allMoves {
			fmt.Println(format.Yellow + "Warning: [--moves] only applies to detailed summary")
			fmt.Println("Usage: pokecli get pokemon --detailed --moves" + format.Reset)
		}
	}
}
