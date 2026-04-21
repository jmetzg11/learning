package main

import "fmt"

// Panics should be used sparinlgy, when there are programmer errors
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
			// recover foo
		}
	}()
	
	f()
}
func f() {
	fmt.Println("a")
	panic("foo")
	fmt.Println("b")
	// a 
}
