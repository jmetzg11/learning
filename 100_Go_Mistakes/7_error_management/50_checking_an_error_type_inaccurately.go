package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		// %w wraps err — outer concrete type becomes *fmt.wrapError,
		// the transientError (if any) is now buried inside the chain
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionID, err)
	}
	return amount, nil
}

func getTransactionAmountFromDB(transactionID string) (float32, error) {
	// ..
	if err != nil {
		return 0, transientError{err: err}
	}
	//..
}

// will always return a 400 for DB errors, even though they're transient
func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")

	amount, err := getTransactionAmount(transactionID)
	if err != nil {
		// LESSON: a type switch checks the concrete TOP-LEVEL type only.
		// It does NOT walk Unwrap() chains. If err is *fmt.wrapError holding
		// a transientError, `case transientError` does not match.
		switch err := err.(type) {
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			// buried transientErrors land here -> 400 instead of 503
			http.Error(w, error.Error(), http.StatusBadRequest)
		}
		return
	}

	// write resposne
}

// solution: errors.As walks the Unwrap() chain and returns true if ANY
// layer matches the target type. So buried transientErrors are found.
func handler(w http.ResponseWriter, r *http.Request) {
	// Get  transaction ID

	amount, err := getTransactionAmount(transactionID)
	if err != nil {
		// (syntax nit: should be &transientError{} not &transientError())
		if errors.As(err, &transientError()) {
			http.Error(w, err.Error(),
				http.StatusServiceUnavailable)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// write resposne
}

// TL;DR
//   type switch / err.(T) → checks concrete top-level type, no unwrap
//   errors.As(err, &T{})  → walks the chain, matches any layer
//   errors.Is(err, target)→ same idea but for comparing to a specific value
//     (e.g. sentinel errors like io.EOF)
