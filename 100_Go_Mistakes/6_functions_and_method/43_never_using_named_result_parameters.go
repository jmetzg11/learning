package main

import "io"

// usually a good idea for interafaces to increase readability.
// debatable on methods.
// Usuage depends on context

func f(a int) (b int) { // name the result to b
	b = a
	return // returns current value of b
}

// what's being returned first, lat or lng?
type locator interface {
	getCoordinates(address string) (float32, float32, error)
}

type locator2 interface {
	getCoordinates(address string) (lat, lng float32, err error)
}

// probably want expressive method signatures too
func (l loc) getCoordinates(address string) (lat, lng float32, err error) {
	// do something
}

// when not to name parameters
func StoreCustomer(customer Customer) (err error) {
	//...
}

// can make implementation short
func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}
