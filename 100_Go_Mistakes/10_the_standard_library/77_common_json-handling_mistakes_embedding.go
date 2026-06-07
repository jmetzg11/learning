package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// MISTAKE: embedding time.Time promotes its MarshalJSON to Event.
// json.Marshal calls time.Time.MarshalJSON — outputs only the timestamp, ID is lost.
type Event1 struct {
	ID int
	time.Time
}

func badEmbedExample() {
	event := Event1{ID: 1234, Time: time.Now()}
	b, _ := json.Marshal(event)
	fmt.Println(string(b)) // "2026-06-06T17:42:46Z" — ID gone
}

// Implementing json.Marshaler lets you control how a type is marshaled.
// MarshalJSON() ([]byte, error) is the full interface — get the name wrong and it's ignored.
type foo struct{}

func (foo) MarshalJSON() ([]byte, error) {
	return []byte(`"foo"`), nil
}

func fooExample() {
	b, _ := json.Marshal(foo{})
	fmt.Println(string(b)) // "foo"
}

// FIX 1: named field instead of embedding — no method promotion, marshals all fields.
type EventFixOne struct {
	ID   int
	Time time.Time
}

// FIX 2: keep embedding but define MarshalJSON to override the promoted one.
type EventFixTwo struct {
	ID int
	time.Time
}

func (e EventFixTwo) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID   int
		Time time.Time
	}{
		ID:   e.ID,
		Time: e.Time,
	})
}

func main() {
	badEmbedExample()
	fooExample()

	b, _ := json.Marshal(EventFixOne{ID: 1234, Time: time.Now()})
	fmt.Println(string(b)) // {"ID":1234,"Time":"2026-06-06T17:42:46Z"}

	b, _ = json.Marshal(EventFixTwo{ID: 1234, Time: time.Now()})
	fmt.Println(string(b)) // {"ID":1234,"Time":"2026-06-06T17:42:46Z"}
}
