package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation struct {
	mu      sync.RWMutex
	balance int
}

// MISTAKE: busy-waiting on a mutex to wait for a condition.
// Each waiter spins, re-checking balance as fast as it can, burning CPU.
func badExample() {
	donation := &Donation{}

	f := func(goal int) {
		donation.mu.RLock()
		// RLock (read) and Lock (write) are mutually exclusive: while a
		// reader holds RLock, the updater's Lock blocks. So the waiter must
		// RUnlock to let the writer in, then RLock again to re-read —
		// that's the wasteful spin.
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			donation.mu.Lock()
			donation.balance++
			stop := donation.balance >= 20
			donation.mu.Unlock()
			if stop {
				return
			}
		}
	}()
}

type DonationChannel struct {
	balance int
	ch      chan int
}

// Channels avoid the busy-wait, but they have their own limitation:
// a value sent on a channel is received by only ONE goroutine, so the
// waiters split the stream instead of all seeing every update. A goal
// can be missed. This is why sync.Cond.Broadcast (wake ALL waiters) fits
// the "notify everyone of a change" problem better than channels.
func betterExample() {
	donation := &DonationChannel{ch: make(chan int)}

	var wg sync.WaitGroup
	wg.Add(2)
	f := func(goal int) {
		defer wg.Done()
		reached := false
		for balance := range donation.ch {
			if !reached && balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				reached = true
			}
		}
	}
	go f(10)
	go f(15)

	for donation.balance < 20 {
		time.Sleep(500 * time.Millisecond)
		donation.balance++
		donation.ch <- donation.balance
	}
	close(donation.ch)
	wg.Wait()
}

type DonationCond struct {
	cond    *sync.Cond
	balance int
}

// FIX: sync.Cond = wait for a condition without busy-waiting, and notify
// every waiter (which a channel can't). A Cond pairs a mutex with a
// "wait until signalled" queue.
//   - cond.Wait() must be called holding the lock. It atomically unlocks,
//     sleeps, and re-locks before returning — so the updater can take the
//     lock and change balance while waiters sleep.
//   - Loop (not if) around Wait: a wake just means "something changed",
//     not "your condition is true", so re-check it.
//   - Broadcast wakes ALL waiters; Signal would wake only one.
func goodExample() {
	donation := &DonationCond{cond: sync.NewCond(&sync.Mutex{})} // cond owns a mutex (cond.L)

	var wg sync.WaitGroup
	wg.Add(2)
	f := func(goal int) {
		defer wg.Done()
		donation.cond.L.Lock()             // must hold the lock to call Wait
		for donation.balance < goal {      // goal not met yet?
			donation.cond.Wait()           // unlock + sleep; on wake, re-lock + loop to re-check
		}
		fmt.Printf("$%d goal reached\n", donation.balance) // condition met, still holding lock
		donation.cond.L.Unlock()
	}
	go f(10) 
	go f(15)

	for {
		time.Sleep(500 * time.Millisecond)
		donation.cond.L.Lock()             // take lock to mutate balance
		donation.balance++
		stop := donation.balance >= 20
		donation.cond.L.Unlock()           // release before notifying
		donation.cond.Broadcast()          // wake BOTH waiters so each re-checks its goal
		if stop {
			break
		}
	}
	wg.Wait()
}
func main() {
	badExample()
	betterExample()
	goodExample()
	time.Sleep(time.Second)
}
