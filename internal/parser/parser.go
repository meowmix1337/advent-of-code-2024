package parser

import (
	"bufio"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	// opens the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// defer will close the file after we have read it into a slice
	defer file.Close()

	// read each line into a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
