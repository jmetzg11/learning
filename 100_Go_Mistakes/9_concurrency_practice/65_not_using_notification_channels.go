package main

import (
	"fmt"
	"time"
)

// Imagine this goroutine is doing a real task: fetching data, writing to disk, etc.
// Once done, it signals the caller via a channel — it has no result to send, just "I'm done".

// BAD: chan bool implies the value matters — true vs false could mean something.
// Does false mean it failed? Does true mean success? Confusing to readers.
func withBool() {
	doneCh := make(chan bool)

	go func() {
		// ... do some work ...
		time.Sleep(10 * time.Millisecond)
		doneCh <- true // why true and not false? ambiguous
	}()

	<-doneCh
	fmt.Println("task finished")
}

// GOOD: chan struct{} makes intent clear — no data, just a "I'm done" ping.
// struct{} also takes zero bytes of memory vs 1 byte for bool.
func withStruct() {
	doneCh := make(chan struct{})

	go func() {
		// ... do some work ...
		time.Sleep(10 * time.Millisecond)
		doneCh <- struct{}{} // clearly just a signal, no value
	}()

	<-doneCh
	fmt.Println("task finished")
}
