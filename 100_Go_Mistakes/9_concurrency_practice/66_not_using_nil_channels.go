package main

// Goal: merge two channels into one.

// BAD: sequential — ch2 won't start until ch1 is fully closed.
// If ch1 is slow, ch2 starves. Not a real merge.
func mergeBad(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for v := range ch1 {
			ch <- v
		}
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

// BROKEN: select interleaves ch1 and ch2 concurrently.
// Problem: when ch1 closes, receiving from a closed channel returns (0, false) immediately.
// The select will keep picking ch1 forever, flooding ch with zero values.
func mergeBroken(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch) // never reached
	}()
	return ch
}

// CLOSER but still wasteful: boolean flags stop forwarding zero values,
// but the closed channel case is still selected on every loop iteration — it
// keeps firing and doing nothing until the other channel also closes.
func mergeInfiniteLoop(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()
	return ch
}

// FIX: set closed channels to nil — a nil channel blocks forever in select,
// completely disabling that case. The for condition exits the loop once both
// are nil, so close(ch) is always reached cleanly.
func mergeGood(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil // select will never pick this case again
					break
				}
				ch <- v
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil // select will never pick this case again
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}
