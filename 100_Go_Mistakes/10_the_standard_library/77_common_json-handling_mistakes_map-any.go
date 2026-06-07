package main

import (
	"encoding/json"
	"fmt"
)

func getMessage() []byte {
	m := map[string]any{
		"id":   32,
		"name": "foo",
	}
	b, _ := json.Marshal(m)
	return b
}

func main() {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}

	// MISTAKE: unmarshaling into map[string]any converts all numbers to float64.
	// The decoder has no type hint, so it defaults to float64 for any JSON number.
	fmt.Printf("%T\n", m["id"]) // float64, not int
}
