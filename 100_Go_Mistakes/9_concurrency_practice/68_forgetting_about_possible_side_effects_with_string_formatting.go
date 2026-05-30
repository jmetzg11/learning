// Be careful about the side effects of string formatting in concurrent applications.
// fmt.Sprintf("%v", x) calls x.String() internally, which READS x's fields.
// If another goroutine is writing to those fields at the same time → data race.

// Data race example (real bug from etcd client)
type Watcher interface {
	// watches on a key or prefix
	// events will be returned through the returned channel
	Watch(ctx contex.Context, key string, opts ...OpOption) WatchChan
	Close() error
}

type watcher struct {
	// streams hold all the active gRPC streams keyed by ctx Value
	streams [mapstring]*watchGrpcStream
}

func (w *watcher) Watch(ctx context.Context, key string, opts ...OpOption) WatchChan {
	// ...
	// DANGER: fmt.Sprintf reads ctx's fields (via ctx.String()).
	// If ctx holds a pointer to a mutable struct and another goroutine
	// modifies that struct concurrently → data race.
	ctxKey := fmt.Sprintf("%v", ctx)
	// ...
	wgs := w.streams[ctxKey] // key provided by the client contained mutable values (e.g. a pointer to a struct)
}

// Solution
// Not to rely on fmt.Sprintf to format the key 
// and create custom streamKeyFromCtx function to extract context value that wasn't mutable 

// Deadlock

type Customer struct {
	mutex sync.RWMutex // protect concurrent accesses 
	id string 
	age int 
}

func (c *Customer) UpdateAgeBad(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}

	c.age = age 
	return nil
}

// Deadlock: UpdateAgeBad holds a write lock, then calls fmt.Errorf("%v", c).
// "%v" triggers c.String(), which tries to acquire a read lock — but the write
// lock is already held by the same goroutine → deadlock.
func (c *Customer) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

// Solution
// Unit test of negative age would catch this. Good argument for unit tests 
// Only perform lock after we confirm the input is valid
func (c *Customer) UpdateAgeGood(age int) error {
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %v", c)
	}
	
	c.mutex.Lock() // only lock after checking if input is valid
	defer c.mutex.Unlock()
	
	c.age = age 
	return nil
}

// Other solution if we can't restrict the scope of a mutex lock
func (c *Customer) UpdateAgeGood2(age int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if age < 0 {
		return fmt.Errorf("age should be positive for customer id %s", c.id) // this does not call .String()
	}

	c.age = age 
	return nil
}