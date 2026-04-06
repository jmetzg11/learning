package main

func main() {

}

// every time bars fills up Go creates another array by doubling its capacity
// also extra effor for the GC to clean all of this up
func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

// no good reason to to give the Go runtime a helping hand
func convert(foos []Foo) []Bar { // options 1
	n := len(foos)
	bars := make([]Bar, 0, n) // initialize with a zero length and give capacity
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}
func convert(foos []Foo) []Bar { // options 2, tends to be slightly faster because it doesn't use append
	n := len(foos)
	bars := make([]Bar, n) // initialize with a give length, because there is a give length all values are 0
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}

// when capacity isn't know at runtime you have a trade of of CPU vs Memory
