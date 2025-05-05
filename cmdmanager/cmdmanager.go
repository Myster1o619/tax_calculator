package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmdManager *CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Enter prices. Confirm each entry with the ENTER key (0 to exit)")
	prices := []string{}
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdManager *CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}