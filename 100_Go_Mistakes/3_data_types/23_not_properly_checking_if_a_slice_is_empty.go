// erroneous
func handleOperations(id string) {
	operations := getOperations(id) // never returns an nil slice
	if operations != nil {
		handle(operations)
	}
}

func getOperations(id string) []float32 {
	operations := make([]float32, 0) // initializes the operations slice

	if id == "" {
		return operations
	}
	// Add elements to operations

	return operations
}

// fix
func handleOperations(id string) {
	operations := getOperations(id) // never returns an nil slice
	if len(operations) != 0 {
		handle(operations)
	}
}

// Go wiki states, when designing interfaces, we should avoid distinguishing nil and empyt slices. Should not make a semantic nor a technical difference if we return a nil or empty slice. Both should mean the same thing for the caller
