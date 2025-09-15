package transformations

import "fmt"

func WeightToString(weight int) string {

	fWeight := float64(weight)

	//convert weight from hectograms to kilograms
	kgWeight := fWeight / 10.00

	//output string in format "[weight] kg"
	kgWeightString := fmt.Sprintf("%.2f", kgWeight)
	return kgWeightString + " kg"
}

func HeightToString(height float64) string {

	fHeight := float64(height)

	//convert height from decimeters to centimeters
	cmHeight := fHeight * 10

	//output string in format "[height] cm"
	cmHeightString := fmt.Sprintf("%2.f", cmHeight)
	return cmHeightString + " cm"
}
