package format

import (
	"fmt"
	"pokecli/api"
	"strings"
)

func PrintPokemonSummary(summary api.PokemonSummary) {
	fmt.Printf("%s (ID: %d)\n", summary.Name, summary.ID)
	fmt.Println("==============================")
	fmt.Printf("Species: %s\n", summary.Species)
	fmt.Printf("Height: %s\n", HeightToString(summary.Height))
	fmt.Printf("Weight: %s\n", WeightToString(summary.Weight))
	fmt.Printf("Types: %s\n", TypesDisplay(summary.Types))
}

func PrintDetailedPokemonSummary(summary api.PokemonSummary, fullMoveset bool) {
	fmt.Printf(Bold+"%s (ID: %d)\n"+Reset, summary.Name, summary.ID)
	fmt.Println(Gray + strings.Repeat("─", 30) + Reset)
	fmt.Printf("Species:   %s\n", summary.Species)
	fmt.Printf("Height:    %s\n", HeightToString(summary.Height))
	fmt.Printf("Weight:    %s\n", WeightToString(summary.Weight))
	fmt.Printf("Types:     %s\n", TypesDisplay(summary.Types))
	fmt.Printf("Abilities: %s\n", AbilitiesDisplay(summary.Abilities))

	fmt.Println("\nStats:")
	for stat, value := range summary.Stats {
		bar := StatBar(value)
		fmt.Printf("  %-15s %3d %s\n", stat, value, bar)
	}

	if len(summary.HeldItems) > 0 {
		fmt.Println("\nHeld Items:")
		for _, item := range summary.HeldItems {
			fmt.Printf("  - %s\n", item)
		}
	}

	fmt.Println("\nMoves:")
	if fullMoveset {
		fmt.Println(MovesDisplay(summary.Moves))
	} else {
		if len(summary.Moves) > 0 {
			for i, move := range summary.Moves {
				if i >= 5 {
					fmt.Println("  ...")
					break
				}
				fmt.Printf("  - %s\n", move)
			}
		}
	}
}
