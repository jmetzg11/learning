package main

import "os"

func main() {
	readFilesError()
	readFile()
	readFilesClosure
}

func readFilesError(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close() // function call when the surrounding function (readFilesError) returns
		// if readFilesError doesn't return then we have a data leak
		// Do something with file
	}

	return nil
}

// Solution 1
func readFiles(ch <-chan string) error {
	for path := range ch {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close() // file is close with each readFile
	// Do somethin with file
	return nil
}

// Solution 2 with closure func
// more complicated to follow but slightly better performace
func readFilesClosure(ch <-chan string) error {
	for path := range ch {
		err := func() error {
			//...
			defer file.Close()
			//...
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
