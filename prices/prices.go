package prices

import (
	"fmt"
	"math"
	"strconv"

	"example.com/tax_calculator/conversion"
	"example.com/tax_calculator/filemanager"
)

type TaxIncludedPrice struct {
	IOManager *filemanager.FileManager `json:"-"` //will be ignored in .json output file
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string][]float64 `json:"tax_included_prices"`
}

func (taxIncludedPrice *TaxIncludedPrice) getPricesFromFile() {
	fileContents, err := taxIncludedPrice.IOManager.ReadLines()

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

	taxIncludedPrice.getPricesFromFile()

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

	// testFileName := fmt.Sprintf("result_%.0f.json", taxIncludedPrice.TaxRate * 100)
	// fmt.Println("testFileName", testFileName)
	// fmt.Println("taxIncludedPrice.TaxRate", taxIncludedPrice.TaxRate)

	err := taxIncludedPrice.IOManager.WriteResult(taxIncludedPrice)

	if err != nil {
		errMsg := fmt.Sprintf("Unable to write contents to JSON file - %v", err)
		fmt.Println(errMsg)
		return
	}
}

func NewTaxIncludedPrice(fm *filemanager.FileManager, taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		IOManager: fm,
		TaxRate: taxRate,
	}
}