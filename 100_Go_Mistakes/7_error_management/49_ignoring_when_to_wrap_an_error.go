// wrap an error to add context or change the error type 


// return the error directly 
if err != nil {
	return err 
}

// old school wrap an error 
type BarError struct {
	Err error 
}

// Satisfies the built-in error interface; called implicitly by fmt.Println, %w, etc.
func (b BarError) Error() string {
	return "bar failed:" +b.Err.Error()
}

// given err = errors.New("db timeout")
if err != nil {
	return BarError{Err: err}
	// fmt.Println -> "bar failed:db timeout"
	// errors.Is(returned, err) -> false  (no Unwrap method)
}

// newer style to wrap error
if err != nil {
	return fmt.Errorf("bar failed: %w", err)
	// fmt.Println -> "bar failed: db timeout"
	// errors.Is(returned, err) -> true   (chain preserved)
}

// transform into another error, source isn't wrapped, but could be part of the error
if err != nil {
	return fmt.Errorf("bar failed: %v", err)
	// fmt.Println -> "bar failed: db timeout"
	// errors.Is(returned, err) -> false  (original err flattened to text)
}

