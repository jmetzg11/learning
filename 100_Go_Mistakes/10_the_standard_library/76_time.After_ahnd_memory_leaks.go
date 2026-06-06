package main

import (
	"context"
	"log"
	"time"
)

type Event struct{}

func handle(_ Event) {}

// MISTAKE: time.After(d) allocates a new timer on every call.
// That timer is not GC'd until d elapses — even if the case is never selected.
// In a busy loop, each iteration leaks a timer that lives for d.
// With d = 1h and frequent events, thousands of live timers accumulate.
func consumer(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour): // new 1-hour timer every iteration
			log.Println("warning: no messages received")
		}
	}
}

// FIX 1: context.WithTimeout — cancel() releases the timer immediately on each event.
func consumerGood(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel() // frees the timer right away
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

// FIX 2: time.NewTimer + Reset — one timer reused forever, no allocation per iteration.
func consumerBetter(ch <-chan Event) {
	timer := time.NewTimer(time.Hour)
	for {
		timer.Reset(time.Hour)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C: // timer.C is a channel; fires (sends) when the duration elapses
			log.Println("warning: no message received")
		}
	}
}
