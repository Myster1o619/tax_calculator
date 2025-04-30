package prices

import (
	"fmt"
	"math"
	"strconv"

	"example.com/tax_calculator/conversion"
	"example.com/tax_calculator/filemanager"
)

type TaxIncludedPrice struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string][]float64
}

func (taxIncludedPrice *TaxIncludedPrice) getPricesFromFile(fileName string) {
	fileContents, err := filemanager.ReadLines(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	convertedFileContent, conversionErr := conversion.StringToFloat(fileContents)

	if conversionErr != nil {
		fmt.Println(conversionErr)
	}
	
	taxIncludedPrice.InputPrices = convertedFileContent
}

func (taxIncludedPrice *TaxIncludedPrice) Process() {

	taxIncludedPrice.getPricesFromFile("prices.txt")

	result := make(map[string][]float64)
	stringKey := strconv.FormatFloat(taxIncludedPrice.TaxRate, 'g', 2, 64)
	adjustedPrices := []float64{}

	for _, price := range taxIncludedPrice.InputPrices {
		adjusted := price * (1 + taxIncludedPrice.TaxRate)
		adjusted = math.Round(adjusted * 100) / 100
		adjustedPrices = append(adjustedPrices, adjusted)
	}

	result[stringKey] = adjustedPrices
	fmt.Println("RESULT:", result)
	taxIncludedPrice.TaxIncludedPrices = result
}

func NewTaxIncludedPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate: taxRate,
	}
}