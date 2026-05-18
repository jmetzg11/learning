//go:build ignore

package main

import (
	"context"
	"net/http"
	"time"
)

// stubs
type Response struct{}

func doSomeTask(ctx context.Context, r *http.Request) (Response, error) { return Response{}, nil }
func publish(ctx context.Context, resp Response) error                  { return nil }
func writeResponse(resp Response)                                       {}

// =============================================================
// BAD - goroutine inherits request ctx, dies when handler returns
// =============================================================
// Timeline:
//   t=0   handler starts, r.Context() is alive
//   t=1   doSomeTask finishes, response ready
//   t=2   goroutine launched with r.Context()
//   t=3   writeResponse returns -> handler returns
//   t=3+  net/http auto-cancels r.Context() (request is done)
//   t=3+  publish() sees ctx.Done() -> aborts mid-flight (Canceled error)
// publish was supposed to keep running AFTER the response went out, but it gets killed.

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := publish(r.Context(), response) // ⚠️ ties async work to request lifetime
		_ = err
	}()

	writeResponse(response)
}

// =============================================================
// GOOD - custom ctx that KEEPS values but DROPS cancellation
// =============================================================
// We still want trace IDs, auth user, locale, etc. (live in ctx.Value)
// but we do NOT want the request's cancellation signal to kill async work.
//
// Trick: wrap the parent ctx in a type that:
//   - returns "no deadline"     -> work has no time budget from parent
//   - returns nil Done channel  -> Done() never fires, so select{<-ctx.Done()} blocks forever
//   - returns nil Err           -> ctx looks "alive" indefinitely
//   - forwards Value lookups    -> request-scoped data still flows through

// the Context interface (for reference - this is what context.Context already is)
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

type detach struct {
	ctx context.Context
}

func (d detach) Deadline() (time.Time, bool) { return time.Time{}, false } // no deadline
func (d detach) Done() <-chan struct{}       { return nil }                // never cancels
func (d detach) Err() error                  { return nil }                // never errored
func (d detach) Value(key any) any           { return d.ctx.Value(key) }   // values still flow

// usage:
func handlerGood(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		// wrap r.Context() in detach -> publish keeps values but ignores parent's cancel
		err := publish(detach{ctx: r.Context()}, response)
		_ = err
	}()

	writeResponse(response)
}

// modern shortcut (Go 1.21+): context.WithoutCancel does exactly the same thing
//   detached := context.WithoutCancel(r.Context())
//   go func() { publish(detached, response) }()
//
// usually pair with a fresh timeout so the async work can't run forever:
//   ctx2, cancel := context.WithTimeout(context.WithoutCancel(r.Context()), 30*time.Second)
//   defer cancel()

