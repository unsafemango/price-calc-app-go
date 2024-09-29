package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

// reading the lines of text from a file
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}
	// bufio is used for reading line by line on a file
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() { // keeps scanning till there is no more content to scan
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err() // find out if an error occured earlier
	if err != nil {
		file.Close()
		return nil, errors.New("failed to read line in file")

	}

	file.Close()
	return lines, nil
}

// write to file in json format
func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create a file")
	}

	encoder := json.NewEncoder(file) // use to convert values to text that follow json format and needs writer type value as an input
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil
}
