// Package format provides CLI rendering functions, text styling utilities, and data transformation helpers for display logic.
package format

import (
	"fmt"
	"pokecli/api"
	"strings"
)

// StatOrder defines the sequence stat values should always be sorted in.
var StatOrder = []string{
	"hp",
	"attack",
	"defense",
	"special-attack",
	"special-defense",
	"speed",
}

// PrintPokemonSummary renders a concise view of Pokémon attributes: name, ID, species, height, weight, and types.
func PrintPokemonSummary(summary api.PokemonSummary) {
	fmt.Printf(Bold+"%s (ID: %d)\n"+Reset, summary.Name, summary.ID)
	fmt.Println(Gray + strings.Repeat("─", 30) + Reset)
	fmt.Printf("Species: %s\n", summary.Species)
	fmt.Printf("Height: %s\n", HeightToString(summary.Height))
	fmt.Printf("Weight: %s\n", WeightToString(summary.Weight))
	fmt.Printf("Types: %s\n", PipeDelimited(summary.Types))
}

// PrintDetailedPokemonSummary renders a detailed view including stats, held items, and moves.
// If fullMoveset is true, all known moves are displayed, otherwise the list is truncated to five.
func PrintDetailedPokemonSummary(summary api.PokemonSummary, fullMoveset bool) {
	PrintPokemonSummary(summary)

	// Stats (Bar Visualizer)
	fmt.Println("\nStats:")
	for _, stat := range StatOrder {
		value := summary.Stats[stat]
		bar, err := StatBar(value)

		if err == nil {
			fmt.Printf("  %-15s %3d %s\n", stat, value, bar)
		} else {
			fmt.Printf("Error: %v", err)
		}
	}

	// Held Items (Bulleted List) if applicable
	if len(summary.HeldItems) > 0 {
		fmt.Println("\nHeld Items:")
		for _, item := range summary.HeldItems {
			fmt.Printf("  - %s\n", item)
		}
	}

	// Moves (Truncated Bulleted List or Full Comma Separated List)
	fmt.Println("\nMoves:")
	if fullMoveset {
		fmt.Println(CommaDelimited(summary.Moves))
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
