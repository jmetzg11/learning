package main

import (
	"database/sql"
	"errors"
)

// Create an error value -> sentinel error
var ErrFoo = errors.New("foo")

func main() {
	err := query()
	if err != nil {
		// will always return false if sql.ErrorNoRows is wrapped in fmt.Errorf with the % w directive
		if err == sql.ErrNoRows {
			// do something from sentinel error
		} else {
			// unexpected error, handle differently
		}
	}

	// Solution
	err := query()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// do something from sentinel error
		} else {
			// unexpected error, handle differently
		}
	}
}

// errors.Is() and error.As() will unwrap the error
