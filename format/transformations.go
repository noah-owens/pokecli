package format

import "fmt"

func WeightToString(weight float64) string {
	//convert weight from hectograms to kilograms
	kgWeight := weight / 10.00

	//output string in format "[weight] kg"
	kgWeightString := fmt.Sprintf("%.2f", kgWeight)
	return kgWeightString + " kg"
}

func HeightToString(height float64) string {
	//convert height from decimeters to centimeters
	cmHeight := height * 10

	//output string in format "[height] cm"
	cmHeightString := fmt.Sprintf("%2.f", cmHeight)
	return cmHeightString + " cm"
}
