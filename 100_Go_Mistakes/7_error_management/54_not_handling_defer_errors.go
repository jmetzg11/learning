package main

import (
	"database/sql"
	"log"
)

const query = "..."

func getBalance1(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	// Use rows
	return 0, nil
}

// better to be explicity with intentions - purposefully ignoring error in defere
func getBalance2(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() { _ = rows.Close() }()

	// Use rows
	return 0, nil
}

// Event better - handle ther error
func getBalance3(db *sql.DB, clientID string) (float32, error) {
	rows, err := db.Query(query, clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		// visibility that the close and query both had errors
		if err != nil {
			if closeErr != nil {
				log.Printf("failed to close rows: %v", closeErr)
			}
			return
		}
		err = closeErr // propagating the error to the caller
	}

	// Use rows
	return 0, nil
}
