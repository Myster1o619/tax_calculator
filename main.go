package main

import (
	"fmt"
	"example.com/tax_calculator/prices"
	"os"
	"strconv"
	"bufio"
)

func getPricesFromFile() ([]float64, error) {
	// dat, _ := os.ReadFile("prices.txt")
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
		return pricesList, err
	}

	return pricesList, nil
	// fmt.Println("dat from getBalanceFromFile", dat)
	// fmt.Println("dat to string",string(dat))
	// convertedString, _ := strconv.ParseFloat(string(dat), 64)
	// return convertedString
}

func main() {
	priceList := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	testIncludedPrices := prices.NewTaxIncludedPrice(0.10)

	dataFromFile, _ := getPricesFromFile()
	fmt.Println(dataFromFile)


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