type Handler struct {
	n int
	publisher publisher
}

type publisher interface {
	Publish([]Foo)
}

func (h Handler) getBestFoo(someInputs int) Foo {
	foos := getFoos(someInputes)
	best := foos[0]

	go runc() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}
		h.publisher.Publish(foos)
	}()
	return best
}


// Test  with sleep
type publisherMock struct {
	mu sync.RWMutex
	got []Foo
}

func (p *publisherMock) Publish(got []Foo) {
	p.mu.Lock()
	defere p.mu.Unlock()
	p.got = got
}

func (p *publisherMock) Get() []Foo {
	p.mu.RLock()
	defere p.mu.RUnlock()
	return p.got
}

func TestGetBestFoo(t *testing.T) {
	mock := publisherMock{}
	h := Handler {
		publisher: &mock,
		n: 2
	}

	foo := h.getBestFoo(42)
	// Check foo

	time.Sleep(10 * time.Millisecond)
	published := mock.Get()
	// Check published
}

// Test with retries — polls until condition is true instead of sleeping a fixed duration
func assert(t *testing.T, assertion func() bool,
	maxRetry int, waitTime time.Duration) {
	for i := 0; i < maxRetry; i++ {
		if assertion() {
			return
		}
		time.Sleep(waitTime)
	}
	t.Fail()
}

func TestGetBestFooRetry(t *testing.T) {
	mock := publisherMock{}
	h := Handler{publisher: &mock, n: 2}

	h.getBestFoo(42)

	// retries up to 30ms total; passes as soon as goroutine finishes
	assert(t, func() bool {
		return len(mock.Get()) == 2
	}, 30, time.Millisecond)
}

// Test with channels
type publisherMock struct {
	ch chan []Foo
}

func (p *publisherMock) Publish(got []Foo) {
	p.ch <- got
}

func TestGetBestFoo(t *testing.T) {
	mock := publisherMock{
		ch: make(chan []Foo),
	}
	defere close(mock.ch)

	h := Handler{
		publisher: &mock,
		n: 2,
	}
	foo := h.getBestFoo(42)
	// Check foo

	if v := len(<-mock.ch); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
}


// Synchronization is better then sleeps, but if it's not possible then retries with small sleeps is better
// than a large passive sleep
