package format

import (
	"fmt"
	"strings"
)

func WeightToString(weight int) string {
	fWeight := float64(weight)

	//convert weight from hectograms to kilograms
	kgWeight := fWeight / 10.00

	//output string in format "[weight] kg"
	kgWeightString := fmt.Sprintf("%.2f", kgWeight)
	return kgWeightString + " kg"
}

func HeightToString(height int) string {
	fHeight := float64(height)

	//convert height from decimeters to centimeters
	cmHeight := fHeight * 10

	//output string in format "[height] cm"
	cmHeightString := fmt.Sprintf("%2.f", cmHeight)
	return cmHeightString + " cm"
}

func StatBar(value int) string {
	barLength := value / 10
	return "[" + strings.Repeat("█", barLength) + strings.Repeat(" ", 10-barLength) + "]"
}

func TypesDisplay(types []string) string {
	return strings.Join(types, " | ")
}

func AbilitiesDisplay(abilities []string) string {
	return strings.Join(abilities, ", ")
}
