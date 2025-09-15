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

	height := format.HeightToString(float64(data.Height))
	weight := format.WeightToString(float64(data.Weight))

	fmt.Printf("Name: %s\nHeight: %s\nWeight: %s\n", data.Name, height, weight)
}
