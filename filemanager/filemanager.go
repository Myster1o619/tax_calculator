package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}

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

func WriteJSON(path string, data interface{}) (error) {
	file, err := os.Create(path)

	if err != nil {
		errorString := fmt.Sprintf("Error creating JSON file %v - %v", path, err)
		err = errors.New(errorString)
		return err
	}

	defer file.Close()

	// jsonData, err := json.Marshal(data)

	// if err != nil {
	// 	errorString := fmt.Sprintf("Error creating JSON data at %v - %v", path, err)
	// 	err = errors.New(errorString)
	// 	return err
	// }

	// _, err = file.Write(jsonData)

	// if err != nil {
	// 	errorString := fmt.Sprintf("Error writing contents to file %v - %v", path, err)
	// 	err = errors.New(errorString)
	// 	return err
	// }

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		errorString := fmt.Sprintf("Error writing JSON encoding to file %v - %v", path, err)
		err = errors.New(errorString)
		return err
	}

	return nil
}