package main

import "fmt"

func main() {
	// don't make any assumptions about ordering with maps
	m := map[string]struct{}{}
	m["a"] = struct{}{}
	m["b"] = struct{}{}
	m["c"] = struct{}{}
	m["d"] = struct{}{}
	m["e"] = struct{}{}
	m["f"] = struct{}{}
	m["g"] = struct{}{}
	for k := range m {
		fmt.Println(k)
	}
	// f
	// g
	// a
	// b
	// c
	// d
	// e

	// map insert during iteration
	l := map[int]bool {
		0: true,
		1: false,
		2: true,
	}
	
	for k, v := range l {
		if v {
			l[10+k] = true
		}
	}
	for k, v := range l {
		fmt.Printf("%s: %+v\n", k, v)
	}
}
