package main

import (
	"fmt"
	"strconv"
)

// A slice is empty if its length is equal to 0
// A slice is nil if it equals nil

func main() {
	var s []string // no allocation
	log(1, s)

	s = []string(nil) // no allocation
	log(2, s)

	s = []string{}
	log(3, s)

	s = make([]string, 0)
	log(4, s)

	// 1: empty=true   nil=true
	// 2: empty=true   nil=true
	// 3: empty=true   nil=false
	// 4: empty=true   nil=false

	var s1 []string // append works on nil slice
	fmt.Println(append(s1, "foo"))
	// [foo]
}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

// favoring nil
// return nil if not foo() or bar()
func f() []string {
	var s []string // favor returning nil, no allocation
	if foo() {
		s = append(s, "foo")
	}
	if bar() {
		s = append(s, "bar")
	}
	return s
}

// favoring empty
func intsToStrings(ints []int) []string {
	s := make([]string, len(ints))
	for i, v := range ints {
		s[i] = strconv.Itoa(v)
	}
	return s
}

// syntactic sugar to pass a nil slice in a single line
s := append([]int(nil), 42)

// create a slice with initial elements
s := []string{"foo", "bar", "baz"}

// Some libraries treat them differently
var s1 []float32  // nil slice
customer1 := customer{
	ID: "foo",
	Operations: s1,
}
b, _ := json.Marshal(customer1)
fmt.Println(string(b))
// {"ID": "foo", "Operations": null}

var s1 := make([]float32, 0) // non-nil, empty slice
customer1 := customer{
	ID: "foo",
	Operations: s1,
}
b, _ := json.Marshal(customer1)
fmt.Println(string(b))
// {"ID": "foo", "Operations": []}
