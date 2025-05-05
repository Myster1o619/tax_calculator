package prices

import (
	"fmt"
	"math"
	"strconv"
	"errors"

	"example.com/tax_calculator/conversion"
	"example.com/tax_calculator/iomanager"
)

type TaxIncludedPrice struct {
	IOManager iomanager.IOManager `json:"-"` //will be ignored in .json output file
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string][]float64 `json:"tax_included_prices"`
}

func (taxIncludedPrice *TaxIncludedPrice) getPricesFromFile() error {
	fileContents, err := taxIncludedPrice.IOManager.ReadLines()

	if err != nil {
		return err
	}

	convertedFileContent, conversionErr := conversion.StringToFloat(fileContents)

	if conversionErr != nil {
		return conversionErr
	}
	
	taxIncludedPrice.InputPrices = convertedFileContent
	return nil
}

func (taxIncludedPrice *TaxIncludedPrice) Process() error {

	err := taxIncludedPrice.getPricesFromFile()

	if err != nil {
		errorString := fmt.Sprintf("Unable to retrieve data - %v", err)
		err = errors.New(errorString)
		return err
	}

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

	return taxIncludedPrice.IOManager.WriteResult(taxIncludedPrice)
}

func NewTaxIncludedPrice(manager iomanager.IOManager, taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		IOManager: manager,
		TaxRate: taxRate,
	}
}