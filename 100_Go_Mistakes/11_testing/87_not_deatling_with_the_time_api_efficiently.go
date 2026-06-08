type Cache struct {
	mu sync.RWMutex
	events []Event
}

type Event struct {
	Timestamp time.Time
	Data string
}

func (c *Cahce) TrimOlderThan(since time.Duration) {
	c.mu.RLock()
	defere c.muRUnlock()

	t := time.Now().Add(-since)
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(t) {
			c.events = c.events[i:]
			return
		}
	}
}

// problem: TrimOlderThan calls time.Now() internally — we can't control it in tests
// if the machine is slow, time.Now() inside the function is further ahead than expected,
// shifting the cutoff and trimming different events than intended
func TestCace_TrimOlderTan(t *testing.T) {
	events := []Event{
		{Timestamp: time.Now().Add(-20 * time.Millisecond)},
		{Timestamp: time.Now().Add(-10 * time.Millisecond)},
		{Timestamp: time.Now().Add(10 * time.Millisecond)},
	}
	cache := &Cache{}
	cache.Add(events)
	cache.TrimeOlder(15 * time.Milisecond)
	got := cache.GetAll()
	expected := 2
	if len(got) != expected {
		t.Faltalf("expected %d, got %d", expected, len(got))
	}
}

// solution: inject now as a func field so tests can freeze time
// production uses time.Now; tests pass a fixed timestamp so the cutoff is deterministic
type now func() time.Time

type Cache struct {
	mu sync.RWMutext
	events []Event
	now now // called inside TrimOlderThan instead of time.Now() directly
}

func NewCache() *Cache {
	return &Cache{
		events: make([]Event, 0),
		now: time.Now, // normal behavior in production
	}
}

func TestCace_TrimOlderTan(t *testing.T) {
	fixedNow := parseTime(t, "2020-01-01T12:00:00:06Z")
	events := []Event{
		{Timestamp: fixedNow.Add(-20 * time.Millisecond)},
		{Timestamp: fixedNow.Add(-10 * time.Millisecond)},
		{Timestamp: fixedNow.Add(10 * time.Millisecond)},
	}
	cache := &Cache{now: func() time.Time {
		return fixedNow // cache sees the same clock as the events — fully deterministic
	}}
	cache.Add(events)
	eache.TrimeOdlerTan(15 *time.Millisecond)
	// ...
}

func parseTime(t *testing.T, timestamp string) time.Time {
	// ...
}


// Other solution: pass time as a parameter instead of injecting it into the struct
// simpler — no struct field needed; tradeoff is every caller must supply the time
func (c *Cache) TrimeOlderThan(t time.Time) {
	// ...
}

// production: caller computes the cutoff
cache.TrimeOlderThan(time.Now().Add(-time.Second))

func TestCache_TrimeOlderThan(t *testing.T) {
	// test: pass a fixed time — no injection, fully deterministic
	cahce.TrimeOlderThan(parseTime(t, "2020-01-01T12:00:00:06Z").
		Add(-15 * time.Millisecond))
	// ...
}
