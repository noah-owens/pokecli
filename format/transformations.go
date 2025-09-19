// Package format provides CLI rendering functions, text styling utilities, and data transformation helpers for display logic.
package format

import (
	"fmt"
	"pokecli/api"
	"strings"
)

// WeightToString converts weight from hg to kg.
// Returns a formatted string in kilograms to two decimal places.
func WeightToString(weight int) string {
	fWeight := float64(weight)

	//convert weight from hectograms to kilograms
	kgWeight := fWeight / 10.00

	kgWeightString := fmt.Sprintf("%.2f", kgWeight)
	return kgWeightString + " kg"
}

// HeightToString converts height from dm to cm.
// Returns a formatted string in centimeters to two decimal places.
func HeightToString(height int) string {
	fHeight := float64(height)

	//convert height from decimeters to centimeters
	cmHeight := fHeight * 10

	cmHeightString := fmt.Sprintf("%2.f", cmHeight)
	return cmHeightString + " cm"
}

// StatBar takes a value and outputs a normalized bar representing its rank between minimum and maximum value
// If successful, returns a bar string and nil error
// If unsuccessful (value is out of bounds) returns empty string and error
func StatBar(value int) (string, error) {
	const minStatValue = 1
	const maxStatValue = 255
	const maxBarLength = 25

	if value < minStatValue || value > maxStatValue {
		return "", fmt.Errorf("stat value %d out of bounds (must be between %d and %d)", value, minStatValue, maxStatValue)
	}

	normalizedBarLength := int(float64(value) / float64(maxStatValue) * float64(maxBarLength))
	bar := "[" + strings.Repeat("█", normalizedBarLength) + strings.Repeat(" ", maxBarLength-normalizedBarLength) + "]"

	return bar, nil
}

// PipeDelimited accepts a slice and converts it to display ready output with a pipe between each term
func PipeDelimited(types []string) string {
	return strings.Join(types, " | ")
}

// CommaDelimited accepts a slice and converts it to a display ready output with a comma between each term
func CommaDelimited(abilities []string) string {
	return strings.Join(abilities, ", ")
}

// PokemonToSummary transforms raw Pokemon API response into display-friendly PokemonSummary.
func PokemonToSummary(p *api.Pokemon) api.PokemonSummary {

	// Extract and transform relevant fields:
	// - Types, Abilities, Moves, HeldItems as slices
	// - Stats as a name-value map

	var types []string
	for _, t := range p.Types {
		types = append(types, t.Type.Name)
	}

	var abilities []string
	for _, a := range p.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}

	stats := make(map[string]int)
	for _, s := range p.Stats {
		stats[s.Stat.Name] = s.BaseStat
	}

	var moves []string
	for _, m := range p.Moves {
		moves = append(moves, m.Move.Name)
	}

	var heldItems []string
	for _, item := range p.HeldItems {
		heldItems = append(heldItems, item.Item.Name)
	}

	return api.PokemonSummary{
		Name:      p.Name,
		ID:        p.ID,
		Height:    p.Height,
		Weight:    p.Weight,
		Types:     types,
		Abilities: abilities,
		Stats:     stats,
		Species:   p.Species.Name,
		Moves:     moves,
		HeldItems: heldItems,
	}
}
