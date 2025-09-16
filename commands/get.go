package commands

import (
	"fmt"
	"os"

	"pokecli/api"
	"pokecli/format"
)

func GetPokemon(args []string) {
	if len(args) != 1 {
		fmt.Println(format.Red + "Usage: pokecli get [pokemon]" + format.Reset)
		os.Exit(1)
	}

	name := args[0]

	data, err := api.FetchPokemon(name)
	if err != nil {
		fmt.Println(format.Red+"Error fetching data: %s"+format.Reset, err)
		os.Exit(1)
	}

	summary := api.PokemonToSummary(data)

	format.PrintSummary(summary)
}
