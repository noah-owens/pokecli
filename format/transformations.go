package format

import (
	"fmt"
	"strings"
)

func WeightToString(weight int) string {
	fWeight := float64(weight)

	//convert weight from hectograms to kilograms
	kgWeight := fWeight / 10.00

	kgWeightString := fmt.Sprintf("%.2f", kgWeight)
	return kgWeightString + " kg"
}

func HeightToString(height int) string {
	fHeight := float64(height)

	//convert height from decimeters to centimeters
	cmHeight := fHeight * 10

	cmHeightString := fmt.Sprintf("%2.f", cmHeight)
	return cmHeightString + " cm"
}

func StatBar(value int) string {
	const maxStatValue = 255
	const maxBarLength = 25

	normalizedBarLength := int(float64(value) / float64(maxStatValue) * float64(maxBarLength))

	return "[" + strings.Repeat("█", normalizedBarLength) + strings.Repeat(" ", maxBarLength-normalizedBarLength) + "]"
}

func TypesDisplay(types []string) string {
	return strings.Join(types, " | ")
}

func AbilitiesDisplay(abilities []string) string {
	return strings.Join(abilities, ", ")
}

func MovesDisplay(moves []string) string {
	return strings.Join(moves, ", ")
}
