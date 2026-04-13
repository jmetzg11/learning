package main

import (
	"fmt"
	"runtime"
)

func main() {
	foos := make([]Foo, 1_000)
	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Foo{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}
// without copying the slice
// 342 KR
// 1024401 KR
// 1024411 KR


// copy the slice 
// 342 KR
// 1024400 KR
// 2453 KR

// marking the remaing variables as nil 
// 342 KR
// 1024401 KR
// 2454 KR

type Foo struct {
	v []byte
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KR\n", m.Alloc/1024)
}

func keepFirstTwoElementsOnly(foos []Foo) []Foo {
	// BAD EXAMPLE
	// return foos[:2]
	
	// Copy the slice 
	res := make([]Foo, 2)
	copy(res, foos)
	return res
	
	// mark the remaing varialbes as nil, maybe if i is closer to n and performace is important, go with this option
	// for i := 2; i < len(foos); i++ {
	// 	foos[i].v = nil
	// }
	// return foos[:2]
}
