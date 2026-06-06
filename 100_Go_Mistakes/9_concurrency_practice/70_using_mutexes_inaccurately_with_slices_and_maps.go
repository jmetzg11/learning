package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

// BAD: lock released before the data is read
func (c *Cache) AverageBalance() float64 {
	c.mu.RLock()
	balances := c.balances // copies the map reference, NOT the data — still a data race
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances { // ranging outside the lock while AddBalance can write
		sum += balance
	}
	return sum / float64(len(balances))
}

// GOOD: hold the lock for the whole read
func (c *Cache) AverageBalanceGood() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

// GOOD (heavy): deep-copy under the lock, then compute lock-free
// use when the computation is slow and you don't want to block writers
func (c *Cache) AverageBalanceGoodHeavy() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}

func main() {
	cache := Cache{balances: make(map[string]float64)}

	for i := 0; i < 4; i++ {
		i := i
		go func() { cache.AddBalance(fmt.Sprintf("user-%d", i), float64(i)*1.11) }()
		go func() { fmt.Println(cache.AverageBalance()) }()
	}

	time.Sleep(100 * time.Millisecond)
}

// Takeaway: a slice/map copied under a lock still shares the underlying data.
// Either hold the lock for the entire read, or deep-copy under the lock.
