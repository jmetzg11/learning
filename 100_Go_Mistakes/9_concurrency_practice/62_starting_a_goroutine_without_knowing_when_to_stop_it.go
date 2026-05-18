//go:build ignore

package main

import "context"

// MISTAKE 62 - starting a goroutine without knowing when to stop it
//
// THE LESSON:
// every `go func() {...}` you spawn needs an answer to:
//   "when and how does this goroutine stop?"
// if you can't answer, you've planted a goroutine LEAK.
// short-running scripts hide leaks; long-running services (web servers,
// daemons) accumulate them over time -> OOM.

// =============================================================
// BAD #1 - goroutine has no stop signal at all
// =============================================================
// foo() hands us a channel. We range over it forever.
// if nobody ever close()s that channel, the goroutine blocks on it
// forever -> permanent leak (memory + scheduler slot).

func foo() <-chan int { return nil } // stub

func memoryLeakExample() {
	ch := foo()
	go func() {
		for v := range ch { // blocks here forever if ch is never closed
			_ = v
			// ...
		}
	}()
}

// =============================================================
// BETTER - pass a ctx so the goroutine CAN be cancelled
// =============================================================
// progress: now the goroutine can react to ctx.Done() and exit.
// remaining problems:
//   - watcher owns resources (file handle, db conn, etc.) - ctx cancel
//     tells it to stop, but does NOT guarantee resources are flushed/closed
//   - lifecycle is implicit: caller has to remember the ctx pattern
//   - cancel() at main-level just means "stop when the whole program ends"

type watcher struct{}

func (w *watcher) watchCtx(ctx context.Context) {
	// select { case <-ctx.Done(): return; case ... }
}

func newWatcherCtx(ctx context.Context) {
	w := &watcher{}
	go w.watchCtx(ctx)
}

func mainCtx() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newWatcherCtx(ctx)
	// Run the application
}

// =============================================================
// GOOD - explicit lifecycle: return a handle with close()
// =============================================================
// caller owns the lifecycle, in plain sight:
//   1. newWatcher returns a *watcher handle
//   2. caller `defer w.close()` -> guaranteed cleanup on exit/panic
//   3. close() signals the goroutine to stop AND releases its resources
//
// big wins:
//   - stop condition is visible right next to where the goroutine starts
//   - close() can FLUSH/CLOSE underlying resources (not just signal)
//   - works for any resource type, not just ctx-aware ones

func (w *watcher) run()   { /* loop until close signal */ }
func (w *watcher) close() { /* signal goroutine to stop + free resources */ }

func newWatcher() *watcher {
	w := &watcher{}
	go w.run()
	return w
}

func mainGood() {
	w := newWatcher()
	defer w.close() // <- guaranteed cleanup on exit (success OR panic)

	// Run the application
}

// =============================================================
// the rule:
//   before writing `go func() { ... }`, answer two questions:
//     1. how does this goroutine learn it's time to stop?
//     2. who is responsible for telling it?
//   if you can't name both, you've planted a leak.
// =============================================================
