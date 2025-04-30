package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	
	if err != nil {
		errorString := fmt.Sprintf("Unable to open file: %v\n", fileName)
		err = errors.New(errorString)
		return nil, err
	}
	
	defer file.Close()
	fileContents := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileContents = append(fileContents, line)
	}

	if err = scanner.Err(); err != nil {
		errorString := fmt.Sprintf("Error reading contents from file: %v - %v\n", fileName, err)
		err = errors.New(errorString)
		return nil, err
	}

	return fileContents, nil
}