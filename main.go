package main

import (
	"fmt"
	"example.com/tax_calculator/prices"
)

func main() {
	priceList := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	testIncludedPrices := prices.NewTaxIncludedPrice(0.10)

	testIncludedPrices.Process()
	fmt.Println("testIncludedPrices.TaxIncludedPrices:", testIncludedPrices.TaxIncludedPrices)

	result := make(map[float64][]float64)

	for _, tax := range taxRates {
		adjustedPrices := []float64{}
		for _, price := range priceList {
			adjusted := price * (1 + tax)
			adjustedPrices = append(adjustedPrices, adjusted)
		}
		result[tax] = adjustedPrices
	}

	for k, v := range result {
		fmt.Println(k,v)
	}
}