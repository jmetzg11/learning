package redis

type Client struct { ... }

type NewClient() *Client { ... }

func (c *Client) Get(key string) (string, error) { ... }

// ==================== //
package client

import "redis"

redis := redis.NewClient() // name collision with package name,
v, err := redis.Get("foo")

// This is allowed but throughout the scope of the redis variable the redis package won't be available

// use a different name

redisClient := redis.NewClient()
v, err := redis.Get("foo")

// play with the import

import redisapi "mylib/redis"

redis := redisapi.NewClient()
v, err := redis.Get("foo")

// also avoid variable names that collide with built in function.

// Rename variables or adjust imports
