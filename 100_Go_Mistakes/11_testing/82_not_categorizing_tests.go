// Build tags
// go test --tags=integration -v .
// 	runs all the tests including integration tests
// go test -v .
// 	runs test with no build tag
//go:build integration

package db

import (
	"os"
	"testing"
)

func TestInsert(t *testing.T) {
	// ...
}

// Environment varialbes
// to avoid skipping tests
func TestInsert(t *testing.T) {
	if os.Getenv("INTEGRATION") != "true" {
		t.Skip("skipping integraiton test")
	}
	// ...
}

// Short mode
// go test -short -v .
func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test")
	}
	// ...
}
