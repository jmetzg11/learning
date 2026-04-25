package main

func f() {
	// ...
	notify() // confusing, did we forget, was this intentional?
}

func notify() error {
	// ...
}

// soltuion for ignoring the error
func g() {
	// worth adding a comment, most cases we don't wanto ignore an error
	_ = notify()
}
