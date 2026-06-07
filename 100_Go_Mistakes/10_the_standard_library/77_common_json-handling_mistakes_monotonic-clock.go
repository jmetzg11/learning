package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Os handles two different clock types: wall and monotonic
// wall clock is subject to variations, shouldn't measure durations with it
// monotonic clock guarantees to move forward

type Event struct {
	Time time.Time
}

func main() {
	t := time.Now()
	event1 := Event{
		Time: t,
	}

	b, err := json.Marshal(event1)
	if err != nil {
		fmt.Println(err)
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(event1 == event2)
	// false
	fmt.Println(event1.Time)
	fmt.Println(event2.Time)
	// 2026-06-06 20:12:28.701680958 -0500 CDT m=+0.000011974
	// 2026-06-06 20:12:28.701680958 -0500 CDT

	// solution 1
	fmt.Println(event1.Time.Equal(event2.Time))
	// true

	solution2()
	// true
}

func solution2() {
	t := time.Now()
	event1 := Event{
		Time: t.Truncate(0),
	}

	b, err := json.Marshal(event1)
	if err != nil {
		fmt.Println(err)
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(event1 == event2)
}
