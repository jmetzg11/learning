package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// BUG: Add is called inside the goroutine, racing with Wait below.
	// Wait can see counter==0 and return before any goroutine runs Add.
	// -> prints 0, 1, 2, or 3 non-deterministically (observed: 0).
	wg := sync.WaitGroup{}
	var v uint64
	for i := 0; i < 3; i++ {
		go func() {
			wg.Add(1)
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(v) // want 3, but races

	// FIX rule: Add must run in the parent, before the goroutine starts.

	// Solution 1: one Add(3) up front (count is known).
	wg1 := sync.WaitGroup{}
	var v1 uint64
	wg1.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg1.Done()
			atomic.AddUint64(&v1, 1)
		}()
	}
	wg1.Wait()
	fmt.Println(v1) // 3

	// Solution 2: Add(1) in the parent each iteration.
	wg2 := sync.WaitGroup{}
	var v2 uint64
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomic.AddUint64(&v2, 1)
		}()
	}
	wg2.Wait()
	fmt.Println(v2) // 3

	// Aside: main does NOT wait for goroutines — it exits as soon as it
	// returns. Without the sleep neither "a" nor "b" prints. Order is still
	// non-deterministic: "ab" or "ba".
	go func() { fmt.Print("a") }()
	go func() { fmt.Print("b") }()
	time.Sleep(10 * time.Millisecond)
}
