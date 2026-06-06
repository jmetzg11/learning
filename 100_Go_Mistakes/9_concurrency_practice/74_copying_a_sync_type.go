package main

import (
	"fmt"
	"sync"
	"time"
)

// MISTAKE: sync types (Mutex, WaitGroup, Cond…) must never be copied.
// A value receiver copies the whole struct — each call gets its own fresh
// mutex, so goroutines never contend for the same lock. No mutual exclusion.

type Counter struct {
	mu       sync.Mutex
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}}
}

// BAD: value receiver copies Counter (and its mutex) on every call
func (c Counter) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++ // modifies the copy, not the original
}

// FIX 1: pointer receiver — all callers share the same mutex and map
func (c *Counter) IncrementGood(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// FIX 2: store the mutex as a pointer — copying the struct is then safe
// because all copies share the same underlying mutex
type CounterGood struct {
	mu       *sync.Mutex
	counters map[string]int
}

func NewCounterGood() CounterGood {
	return CounterGood{
		mu:       &sync.Mutex{},
		counters: map[string]int{},
	}
}

func (c CounterGood) Increment(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := NewCounter()
	go func() { c.IncrementGood("foo") }()
	go func() { c.IncrementGood("bar") }()
	time.Sleep(10 * time.Millisecond)
	fmt.Println(c.counters) // map[bar:1 foo:1]
}
