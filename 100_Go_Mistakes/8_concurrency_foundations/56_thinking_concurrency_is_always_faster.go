//go:build ignore

package main

import "sync"

// A thread is the smallest unit of processing that an OS can perform.
//
// concurrent - two or more threads can start, run, and complete in
//              overlapping time periods
// parallel   - the same task can be executed multiple times at once
//
// CPU switching from one thread to another = context switching.
// thread moves from executing -> runnable state
//
// goroutines = application-level threads
// goroutine context switch is 80-90% faster than CPU thread switch
//
//   G -> Goroutine
//   M -> OS thread (machine)
//   P -> CPU core (processor)
//
// Each M is assigned to a P by the OS scheduler. Each G runs on an M.
// GOMAXPROCS = limit of Ms executing user-level Go code simultaneously.
// Default = number of available CPU cores.
//
// Goroutine states: Executing | Runnable | Waiting
//
// When all Ps are busy, the runtime queues goroutines.
// Each P has a local queue; there's also a global queue.

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

func sequentialMergesort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle])
	sequentialMergesort(s[middle:])
	merge(s, middle)
}

// SLOWER than sequential!
// Goroutine creation + scheduling cost dominates the work of merging
// a tiny number of items.
func parallelMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)
}

// SOLUTION
// pick threshold via benchmark on a prod-like machine.
// below it, stay sequential to amortize goroutine cost.
const max = 2048

func parallelMergesortV2(s []int) {
	if len(s) <= 1 {
		return
	}

	if len(s) <= max {
		sequentialMergesort(s)
	} else {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}
