package conversion

import(
	"strconv"
	"errors"
	"fmt"
)

func StringToFloat(strings []string) ([]float64, error) {
	floatList := make([]float64, len(strings))

	for _, stringValue := range strings {
		convertedLine, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			conversionError := fmt.Sprintf("Error converting %v to float - %v\n", stringValue, err)
			err = errors.New(conversionError)
			return nil, err
		}

		floatList = append(floatList, convertedLine)
	}

	return floatList, nil
}