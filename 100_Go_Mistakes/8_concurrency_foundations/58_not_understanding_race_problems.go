//go:build ignore

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Data Race - two or more goroutines simultaneously access the same
// memory location, and at least one is writing.

func dataRace() {
	i := 0
	go func() { i++ }()
	go func() { i++ }()
	_ = i
}

// `go run -race` / `go test -race` will warn of race conditions.

// Fix #1: sync/atomic (works for ints, pointers; NOT slices/maps/structs)
func atomicFix() {
	var i int64 = 0
	go func() { atomic.AddInt64(&i, 1) }()
	go func() { atomic.AddInt64(&i, 1) }()
}

// Fix #2: Mutex (mutual exclusion)
func mutexFix() {
	var mu sync.Mutex
	i := 0
	go func() {
		mu.Lock()
		i++
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		i++
		mu.Unlock()
	}()
	_ = i
}

// Fix #3: communicate across goroutines via channel
func channelFix() {
	i := 0
	ch := make(chan int)
	go func() { ch <- 1 }()
	go func() { ch <- 1 }()
	i += <-ch
	i += <-ch
	fmt.Println(i)
}

// Race Condition (different from data race):
// depends on execution order, not simultaneous memory access.
// Non-deterministic results.
func raceCondition() {
	var mu sync.Mutex
	i := 0
	go func() {
		mu.Lock()
		defer mu.Unlock()
		i = 1
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		i = 2
	}()
	_ = i
	// safe (no data race) but result is 1 or 2 depending on order
}

// THE GO MEMORY MODEL

// no race - i++ runs in goroutine, nothing reads i in main
func noRace() {
	i := 0
	go func() { i++ }()
	_ = i
}

// RACE - main reads i while goroutine writes it
func race() {
	i := 0
	go func() { i++ }()
	fmt.Println(i)
}

// no race - channel send/receive is a sync point (happens-before)
func noRaceUnbuffered() {
	i := 0
	ch := make(chan struct{}) // unbuffered: send blocks until receive
	go func() {
		<-ch              // blocks here until main sends
		fmt.Println(i)    // prints 1 (i++ happened-before this)
	}()
	i++                   // i = 1
	ch <- struct{}{}      // rendezvous: unblocks the goroutine
}

// variable increment < channel send < channel receive < variable read

// no race - close(ch) is also a sync point (happens-before any receive)
func noRaceClose() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch              // receive on closed chan returns zero value immediately
		fmt.Println(i)    // prints 1
	}()
	i++                   // i = 1
	close(ch)             // unblocks goroutine without sending a value
}

// RACE - buffered channel doesn't sync this way
func raceBuffered() {
	i := 0
	ch := make(chan struct{}, 1) // buffer size 1
	go func() {
		i = 1
		<-ch                      // receive happens AFTER main's send returns
	}()
	ch <- struct{}{}              // doesn't block (buffer has room) - no rendezvous
	fmt.Println(i)                // may run before goroutine's i = 1
}

// no race - unbuffered send blocks until receive (rendezvous)
func noRaceRendezvous() {
	i := 0
	ch := make(chan struct{})
	go func() {
		i = 1                     // happens-before the receive below
		<-ch
	}()
	ch <- struct{}{}              // blocks until goroutine's <-ch runs
	fmt.Println(i)                // prints 1 - i = 1 is guaranteed visible
}
