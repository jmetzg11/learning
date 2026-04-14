package main

import (
	"fmt"
	"runtime"
)

// The number of buckets in a map cannot shrink, think of a cache -> it would be a memory leak
func main() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)

}

// 342 KR
// 297709 KR
// 294689 KR
func randBytes() [128]byte {
	return [128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KR\n", m.Alloc/1024)
}

// Solution 1
// Create a copy of the map and discard the old map. Problems is memory consumption will be twice as high for brief period

// Solution 2
// map stores an array pointer, GC will get the pointers and the buckets will hold 8bytes instead of an array of 128
// map[int]*[128]byte
