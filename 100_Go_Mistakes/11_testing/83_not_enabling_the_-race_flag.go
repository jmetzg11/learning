// go test -race ./...
// memory usage may increase by 5 to 10x 
// execution time may increase by 2 to 20x 

// cannot give false positives, but false negatives are possible 
// to increase chances of catching all race conditions put tests in loops 
func testDataRace(t *testing.T) {
	for i := 0; i < 100; i++ {
		// Actual logic
	}
}

// if a specific file contains test that lead to data races
// we can exclud it from the race detection using the !race build tag 

//go:build !race 

package main 

import (
	"testing"
)

func TestFoo(t *testing.T) {
	// ...
}

func TestBar(t *testing.T) {
	// ...
}

// file will be skipped if the race flag is enabled 