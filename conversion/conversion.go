package conversion

import(
	"strconv"
	"errors"
	"fmt"
)

func StringToFloat(strings []string) ([]float64, error) {
	floatList := make([]float64, len(strings))
	// floatList := []float64{}

	for index := 0; index < len(strings); index++ {
		convertedLine, err := strconv.ParseFloat(strings[index], 64)

		if err != nil {
			conversionError := fmt.Sprintf("Error converting %v to float - %v\n", strings[index], err)
			err = errors.New(conversionError)
			return nil, err
		}

		// floatList = append(floatList, convertedLine)
		floatList[index] = convertedLine
	}

	// for _, stringValue := range strings {
	// 	convertedLine, err := strconv.ParseFloat(stringValue, 64)

	// 	if err != nil {
	// 		conversionError := fmt.Sprintf("Error converting %v to float - %v\n", stringValue, err)
	// 		err = errors.New(conversionError)
	// 		return nil, err
	// 	}

	// 	floatList = append(floatList, convertedLine)
	// }

	return floatList, nil
}