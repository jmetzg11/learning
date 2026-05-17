//go:build ignore

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// A Context carries a deadline, a cancellation signal, and other
// request-scoped values across API boundaries.

// =============================================================
// DEADLINE
// =============================================================
// A specific point in time:
//   time.Duration   from now
//   time.Time       absolute

// stubs for the example (would normally come from flight + kafka pkgs)
type Position struct {
	Lat, Lng float64
}

type kafkaClient struct{}

func (k *kafkaClient) Send(p Position) error { return nil }

// --- the interface: contract that "a publisher accepts a ctx" ---
type publisher interface {
	Publish(ctx context.Context, position Position) error
}

// --- impl #1: talks to Kafka ---
type kafkaPublisher struct {
	client *kafkaClient
}

func (k kafkaPublisher) Publish(ctx context.Context, p Position) error {
	done := make(chan error, 1)
	go func() {
		done <- k.client.Send(p) // imagine this could hang for minutes
	}()

	select {
	case err := <-done:
		return err // work finished in time
	case <-ctx.Done():
		return ctx.Err() // 4s elapsed -> DeadlineExceeded, give up
	}
}

// --- impl #2: POSTs to an HTTP API ---
type httpPublisher struct {
	url    string
	client *http.Client
}

func (h httpPublisher) Publish(ctx context.Context, p Position) error {
	body, _ := json.Marshal(p)

	// attach ctx to the request -> net/http will abort the call if ctx cancels
	req, err := http.NewRequestWithContext(ctx, "POST", h.url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	// if the server takes > 4s, Do() returns ctx.Err() automatically.
	// no manual select needed - the stdlib watches ctx.Done() for us.
	resp, err := h.client.Do(req)
	if err != nil {
		return err // could be DeadlineExceeded if timeout fired mid-call
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("publish failed: %s", resp.Status)
	}
	return nil
}

// --- handler holds ANY publisher (real, fake, http, kafka...) ---
type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position Position) error {
	// Background() = root ctx (never cancels on its own)
	// WithTimeout wraps it: auto-cancels after 4s
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel() // free the timer even if it already fired

	return h.pub.Publish(ctx, position) // 4s budget handed to Publish
}

// timeline if broker is slow:
//   t=0s   Publish starts with 4s budget
//   t=4s   ctx auto-cancels -> ctx.Done() channel closes
//          select picks the <-ctx.Done() case -> returns DeadlineExceeded
//          caller fails fast instead of hanging forever

// =============================================================
// CANCELLATION SIGNALS
// =============================================================
// WithCancel != WithTimeout - it DOESN'T auto-cancel.
// It cancels ONLY when you (or a parent ctx) call cancel().
// Use it for "stop when I say so" instead of "stop after Xs".

func CreateFileWatcher(ctx context.Context, name string) {} // stub

func cancellationExample() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // <- THIS is the signal:
	//    when this func returns, cancel() fires
	//    -> ctx.Done() closes
	//    -> goroutine sees it and cleans up

	go func() {
		CreateFileWatcher(ctx, "foo.txt") // watcher must select on ctx.Done() to quit
	}()
}

// inside the watcher (sketch):
//   for {
//       select {
//       case <-ctx.Done():        // parent said stop -> exit, free resources
//           return
//       case ev := <-fileEvents:  // do normal work
//           handle(ev)
//       }
//   }

// the three "WithX" flavors:
//   WithTimeout(parent, 4s)    -> auto-cancels after 4s
//   WithDeadline(parent, t)    -> auto-cancels at absolute time t
//   WithCancel(parent)         -> only cancels when YOU call cancel()
// all three return a cancel func - ALWAYS defer it (avoids leaks)

// =============================================================
// CONTEXT VALUES
// =============================================================
// ctx can also carry per-request data (auth user, trace id, feature flags...)
// rule of thumb: only use for REQUEST-SCOPED values that cross API boundaries
// NOT for passing normal function args - that's an anti-pattern

// custom unexported type prevents key collisions across packages.
// if everyone used "userID" as a string key, packages would overwrite each other.
type key string

const isValidHostKey key = "isValidHost"

// HTTP middleware: wraps next handler, injects a value, hands ctx down
func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validHost := r.Host == "acme"

		// WithValue returns a NEW child ctx carrying the key/value pair.
		// parent ctx is untouched - context is immutable, only nesting grows it
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)

		// attach new ctx to the request, pass to next handler in chain
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// downstream handler reads it:
func nextHandler(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(isValidHostKey) // returns interface{}
	valid, ok := v.(bool)                  // type assert back to bool
	if !ok || !valid {
		http.Error(w, "bad host", 403)
		return
	}
	// ... handle request
}

// flow:
//   request -> checkValid middleware -> adds isValidHost -> nextHandler reads it
// key insights:
//   - context.WithValue is a linked list of {key, value, parent} - O(n) lookup
//   - use ctx values sparingly; prefer explicit function args when possible
//   - good fits: trace IDs, auth principal, request ID, locale


// =============================================================
// Catching a context cancellation
// =============================================================
// Standard "long-running worker" pattern: a goroutine loops forever,
// processing messages AND watching for cancellation in the same select.
//
// Without the <-ctx.Done() case, the loop could block on ch forever
// even after the caller signaled "give up".

type Message struct{} // stub

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		case msg := <-ch:
			_ = msg
			// Do something with msg
		case <-ctx.Done():
			return ctx.Err() // tells caller WHY: Canceled or DeadlineExceeded
		}
	}
}

// =============================================================
// Implementing a function that receives a context
// =============================================================
// Rule: if your function takes a ctx, every blocking op inside it
// must respect ctx. Otherwise the ctx is a LIE - the caller thinks
// they can cancel, but the function will still hang.

// stubs used by fBad / fGood below
var (
	ch1 = make(chan struct{})
	ch2 = make(chan int)
)

// BAD - blind channel ops ignore ctx
//   if no one is reading ch1, the send blocks forever
//   if no one ever sends to ch2, the receive blocks forever
//   ctx.Done() firing changes nothing -> caller can't actually cancel
func fBad(ctx context.Context) error {
	// ...
	ch1 <- struct{}{}

	v := <-ch2
	_ = v
	// ...
	return nil
}

// GOOD - every channel op is in a select with ctx.Done()
//   now cancellation always takes effect within microseconds
func fGood(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case ch1 <- struct{}{}:
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case v := <-ch2:
		_ = v
		// ...
	}
	return nil
}

// =============================================================
// context.TODO()
// =============================================================
// Behaves identically to Background() at runtime - both return an
// empty, never-cancels ctx.
//
// The difference is INTENT, picked up by linters and code review:
//   Background()  -> "this IS the root ctx (main, init, tests)"
//   TODO()        -> "I don't have a real ctx here YET, fix later"
//
// Use TODO when adding ctx to legacy code incrementally - it marks
// spots where a real ctx should eventually be plumbed through.

func someAPIThatNeedsCtx(ctx context.Context, key string) (string, error) {
	return "", nil
}

func legacyFunc() {
	// no ctx available here (yet) - use TODO as a placeholder
	// reviewers/tools can grep for context.TODO() to find work to do
	result, err := someAPIThatNeedsCtx(context.TODO(), "foo")
	_, _ = result, err
}

// once legacyFunc is updated to accept a ctx, the TODO goes away:
func legacyFuncFixed(ctx context.Context) {
	result, err := someAPIThatNeedsCtx(ctx, "foo")
	_, _ = result, err
}
