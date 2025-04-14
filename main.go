package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	result := make(map[float64][]float64)

	for _, tax := range taxRates {
		var adjustedPrices []float64
		for _, price := range prices {
			adjusted := price - (price * tax)
			adjustedPrices = append(adjustedPrices, adjusted)
		}
		result[tax] = adjustedPrices
	}

	for k, v := range result {
		fmt.Println(k,v)
	}
}