package format

import (
	"fmt"
	"pokecli/api"
	"strings"
)

func PrintSummary(summary api.PokemonSummary) {
	fmt.Printf("%s (ID: %d)\n", summary.Name, summary.ID)
	fmt.Println("==============================")
	fmt.Printf("Species: %s\n", summary.Species)
	fmt.Printf("Height: %s\n", HeightToString(summary.Height))
	fmt.Printf("Weight: %s\n", WeightToString(summary.Weight))
	fmt.Printf("Types: %s\n", strings.Join(summary.Types, ", "))
	fmt.Printf("Abilities: %s\n", strings.Join(summary.Abilities, ", "))
	fmt.Println("Stats:")
	for stat, value := range summary.Stats {
		fmt.Printf("  %s: %d\n", stat, value)
	}
}
