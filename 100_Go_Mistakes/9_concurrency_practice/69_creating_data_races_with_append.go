package main

import (
	"fmt"
	"time"
)

func main() {
	s := make([]int, 1)

	go func() {
		s1 := append(s, 1)
		fmt.Println(s1)
	}()

	go func() {
		s2 := append(s, 1)
		fmt.Println(s2)
	}()

	// NO race: len==cap so append allocates a new array per goroutine
	time.Sleep(100 * time.Millisecond)
	
	y := make([]int, 0, 1)

	go func() {
		y1 := append(y, 1)
		fmt.Println(y1)
	}()

	go func() {
		y2 := append(y, 1)
		fmt.Println(y2)
	}()
	
	// WARNING: DATA RACE — len<cap so both goroutines write to the same backing array
	// Write at 0x00c000188048 by goroutine 11:
	time.Sleep(100 * time.Millisecond)

	x := make([]int, 0, 1)

	go func() {
		xCopy := make([]int, len(x), cap(x))
		copy(xCopy, x)

		x1 := append(xCopy, 1)
		fmt.Println(x1)
	}()

	go func() {
		xCopy := make([]int, len(x), cap(x))
		copy(xCopy, x)

		x2 := append(xCopy, 1)
		fmt.Println(x2)
	}()

	time.Sleep(100 * time.Millisecond)
}

// Slice race rule: safe when len==cap (append allocates new array), race when len<cap (shared backing array)
// Fix: each goroutine copies the slice before appending (see x example above)

// Maps: concurrent reads+writes are always a data race, even on different keys
// Fix: use a mutex or sync.Map 