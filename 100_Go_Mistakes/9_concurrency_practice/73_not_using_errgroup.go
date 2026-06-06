package main

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Circle struct{ r float64 }
type Result struct{ area float64 }

func foo(ctx context.Context, c Circle) (Result, error) {
	if c.r < 0 {
		return Result{}, fmt.Errorf("negative radius")
	}
	return Result{c.r * c.r * 3.14}, nil
}

// MISTAKE: WaitGroup gives no way to return errors from goroutines.
// The only options (shared slice, channel, single var) all require extra
// synchronization and still can't cancel the other goroutines on failure.
func badHandler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	var firstErr error
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(len(circles))

	for i, circle := range circles {
		i := i
		circle := circle
		go func() {
			defer wg.Done()
			result, err := foo(ctx, circle)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err // capturing the first error requires its own mutex
				}
				mu.Unlock()
				return
			}
			results[i] = result
		}()
	}

	wg.Wait()
	return results, firstErr
}

// FIX: errgroup = WaitGroup + error collection + context cancellation in one.
// Goroutines return error directly; g.Wait() returns the first non-nil error
// and the derived context is cancelled automatically on any failure.
func goodHandler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	return results, g.Wait()
}

func main() {
	circles := []Circle{{1}, {2}, {-1}, {4}}
	ctx := context.Background()

	_, err := badHandler(ctx, circles)
	fmt.Println("bad:", err)

	_, err = goodHandler(ctx, circles)
	fmt.Println("good:", err)
}
