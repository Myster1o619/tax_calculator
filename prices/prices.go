package prices

import (
	"fmt"
	"strconv"
	"math"
	"os"
	"bufio"
)

type TaxIncludedPrice struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string][]float64
}

func (taxIncludedPrice *TaxIncludedPrice) getPricesFromFile() {
	file, _ := os.Open("prices.txt")
	defer file.Close()
	pricesList := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		convertedLine, _ := strconv.ParseFloat(line, 64)
		pricesList = append(pricesList, convertedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return 
	}

	taxIncludedPrice.InputPrices = pricesList
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
}

func NewTaxIncludedPrice(taxRate float64) *TaxIncludedPrice {
	return &TaxIncludedPrice{
		TaxRate: taxRate,
	}
}