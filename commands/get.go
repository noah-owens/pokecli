package commands

import (
	"fmt"
	"os"

	"pokecli/api"
	"pokecli/format"
)

func Get(args []string) {
	if len(args) != 1 {
		fmt.Println(format.Red + "Usage: pokecli get [pokemon]" + format.Reset)
		os.Exit(1)
	}

	name := args[0]
	data, err := api.FetchPokemon(name)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", data.Name, data.Height, data.Weight)
}
