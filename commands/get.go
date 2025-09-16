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
		fmt.Println(format.Red + "Usage: pokecli get pokemon [name] [--detailed]" + format.Reset)
		os.Exit(1)
	}

	name := args[0]
	detailed := false

	detailedFlags := []string{"--detailed", "-d"}

	for _, arg := range args[1:] {
		if slices.Contains(detailedFlags, arg) {
			detailed = true
		}
	}

	data, err := api.FetchPokemon(name)
	if err != nil {
		fmt.Printf(format.Red+"Error fetching data: %s"+format.Reset, err)
		os.Exit(1)
	}

	summary := api.PokemonToSummary(data)

	if detailed {
		format.PrintDetailedPokemonSummary(summary)
	} else {
		format.PrintPokemonSummary(summary)
	}
}
