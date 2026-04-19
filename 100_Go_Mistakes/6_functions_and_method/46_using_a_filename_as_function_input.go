package main

import (
	"bufio"
	"io"
	"os"
)

// How to unit test this? Need to make files for the test
// Not reusable, can't use it for other cases for counting lines e.g. HTTP request
func countEmptylinesInFileBad(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// Handle file closure

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // iterates over each line
		// ...
	}
	return 10
}

// Solution
// reuse regardless of input type
// easier to test without so many different files
func countEmptyLines(reader io.Reader) (int, err) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		//...
	}
}
