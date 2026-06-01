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

func (c *Cache) AverageBalanceGood() float64 {

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

// Fix: either hold the lock for the entire loop, or deep copy the map under the lock:
//   c.mu.RLock()
//   balances := make(map[string]float64, len(c.balances))
//   for k, v := range c.balances { balances[k] = v }
//   c.mu.RUnlock()
