package main

import (
	"fmt"

	// "example.com/tax_calculator/cmdmanager"
	"example.com/tax_calculator/filemanager"
	"example.com/tax_calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.10, 0.15}

	for _, rate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", rate * 100))
		// cmdManager := cmdmanager.New()
		taxIncludedPriceJob := prices.NewTaxIncludedPrice(fileManager, rate)
		taxIncludedPriceJob.Process()
	}

	// result := make(map[float64][]float64)

	// for _, tax := range taxRates {
	// 	adjustedPrices := []float64{}
	// 	for _, price := range testIncludedPrices.InputPrices {
	// 		adjusted := price * (1 + tax)
	// 		adjustedPrices = append(adjustedPrices, adjusted)
	// 	}
	// 	result[tax] = adjustedPrices
	// }

	// for k, v := range result {
	// 	fmt.Println(k,v)
	// }
}