package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// reading the lines of text from a file
func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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
func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

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

func New(inputpath string, outputpath string) FileManager {
	return FileManager{
		InputFilePath:  inputpath,
		OutputFilePath: outputpath,
	}
}
