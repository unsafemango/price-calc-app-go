package filemanager

import (
	"bufio"
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
