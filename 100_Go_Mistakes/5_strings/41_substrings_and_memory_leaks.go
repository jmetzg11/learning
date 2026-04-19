package main

import "fmt"

func main() {
	s1 := "Hêllo, World!"
	s2 := s1[:5] // gets the first 5 bytes - not runes! ,
	fmt.Println(s2)
	// Hêll

	s3 := string([]rune(s1)[:5])
	fmt.Println(s3)
	// Hêllo
}

type store struct{}

func (s store) store(uuid string) {
	//...
}

func (s store) handleLog1(log string) error {
	if len(log) < 36 {
		return errors.New "log is not correctly formatted"
	}
	uuid := log[:36] // referencing the same backing array, including the bytes in the initial log string
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog2(log string) error {
	if len(log) < 36 {
		return errors.New "log is not correctly formatted"
	}
	uuid := string([]byte(log[:36])) // the backing memory array has just has 36 bytes
	s.store(uuid)
	// Do something
	return nil
}

func (s store) handleLog3(log string) error {
	if len(log) < 36 {
		return errors.New "log is not correctly formatted"
	}
	uuid := strings.Clone(log[:36])
	s.store(uuid)
	// Do something
	return nil
}
