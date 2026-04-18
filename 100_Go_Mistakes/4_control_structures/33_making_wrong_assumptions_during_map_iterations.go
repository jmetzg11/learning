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
	l := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	// this is unpredictable, may be produced during the interaction or skipped.
	for k, v := range l {
		if v {
			l[10+k] = true
		}
	}
	fmt.Println(l)
	// couple of possible outputs
	// map[0:true 1:false 2:true 10:true 12:true 20:true 22:true]}
	// map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true 40:true]}

	// solution -> make a copy of the map
	k := map[int]bool{
		0: true,
		1: false,
		2: true,
	}
	k2 := copyMap(k)
	for k, v := range k {
		k2[k] = v
		if v {
			k2[10+k] = true
		}
	}
	fmt.Println(k2)
	// map[0:true 1:false 2:true 10:true 12:true]
}

func copyMap(m map[int]bool) map[int]bool {
	res := make(map[int]bool, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}
