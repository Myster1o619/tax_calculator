package prices

import (
	"fmt"
	"strconv"
)

type TaxIncludedPrice struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string][]float64
}

func (taxIncludedPrice *TaxIncludedPrice) Process() {
	result := make(map[string][]float64)
	stringKey := strconv.FormatFloat(taxIncludedPrice.TaxRate, 'g', 2, 64)
	adjustedPrices := []float64{}

	for _, price := range taxIncludedPrice.InputPrices {
		adjusted := price * (1 + taxIncludedPrice.TaxRate)
		adjustedPrices = append(adjustedPrices, adjusted)
	}

	result[stringKey] = adjustedPrices
	fmt.Println("RESULT:", result)
	taxIncludedPrice.TaxIncludedPrices = result
}

func NewTaxIncludedPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate: taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}