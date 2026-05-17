package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// =============================================================
// THE CONTRACT (the interface)
// =============================================================
// "Anything with a Publish(ctx, Position) error method is a publisher."
// publishHandler will accept ANY of these without knowing which.

type Position struct {
	Lat, Lng float64
}

type publisher interface {
	Publish(ctx context.Context, p Position) error
}

// =============================================================
// IMPL #1 - just prints (zero deps, instant)
// =============================================================
type consolePublisher struct {
	name string
}

func (c consolePublisher) Publish(ctx context.Context, p Position) error {
	fmt.Printf("[console:%s] %+v\n", c.name, p)
	return nil
}

// =============================================================
// IMPL #2 - real HTTP POST
// =============================================================
type httpPublisher struct {
	url    string
	client *http.Client
}

func (h httpPublisher) Publish(ctx context.Context, p Position) error {
	body, _ := json.Marshal(p)

	// attach ctx to the request -> http client auto-aborts if ctx cancels
	req, err := http.NewRequestWithContext(ctx, "POST", h.url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return err // could be DeadlineExceeded if timeout fired mid-request
	}
	defer resp.Body.Close()

	fmt.Printf("[http] POST %s -> %s\n", h.url, resp.Status)
	return nil
}

// =============================================================
// IMPL #3 - slow, designed to trigger the 4s timeout
// =============================================================
type slowPublisher struct{}

func (s slowPublisher) Publish(ctx context.Context, p Position) error {
	done := make(chan struct{})
	go func() {
		time.Sleep(6 * time.Second) // simulate slow downstream
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("[slow] eventually published")
		return nil
	case <-ctx.Done():
		return ctx.Err() // ctx hits 4s first -> DeadlineExceeded
	}
}

// =============================================================
// THE HANDLER - knows NOTHING about which impl it holds
// =============================================================
type publishHandler struct {
	pub publisher // <-- holds any type satisfying the interface
}

func (h publishHandler) publishPosition(p Position) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	return h.pub.Publish(ctx, p) // same call, different code runs underneath
}

// =============================================================
// MAIN - swap implementations through the SAME handler
// =============================================================
func main() {
	pos := Position{Lat: 37.77, Lng: -122.42}

	publishers := []publisher{
		consolePublisher{name: "stdout"},
		httpPublisher{url: "https://httpbin.org/post", client: &http.Client{}},
		slowPublisher{},
	}

	for _, p := range publishers {
		h := publishHandler{pub: p} // ONE handler shape, different innards
		start := time.Now()
		err := h.publishPosition(pos)
		fmt.Printf("  -> err=%v elapsed=%v\n\n",
			err, time.Since(start).Round(time.Millisecond))
	}
}

// Expected output:
//
// [console:stdout] {Lat:37.77 Lng:-122.42}
//   -> err=<nil> elapsed=0s
//
// [http] POST https://httpbin.org/post -> 200 OK
//   -> err=<nil> elapsed=~500ms
//
//   -> err=context deadline exceeded elapsed=4s
//
// Notice: publishHandler.publishPosition() has IDENTICAL code in all three
// calls. Only the underlying type differs. That's the whole point of an
// interface - the caller is decoupled from the implementation.
