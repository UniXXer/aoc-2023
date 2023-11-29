package commons

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readerToArray(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func readFileIntoArray(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	inputs, err := readerToArray(file)

	if err != nil {
		return nil, err
	}

	return inputs, nil
}

func ReadExample(day string) ([]string, error) {
	var filename string = fmt.Sprintf("day%s/inputs/example.txt", day)
	return readFileIntoArray(filename)
}

func ReadInput(day string) ([]string, error) {
	var filename string = fmt.Sprintf("day%s/inputs/input.txt", day)
	return readFileIntoArray(filename)
}
