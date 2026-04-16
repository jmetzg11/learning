package main

import "fmt"

func main() {
	s := []int{0, 1, 2}
	t := make([]int, len(s))
	copy(s, t)

	for range s { // s gets a slice copy
		s = append(s, 10)
	}
	fmt.Println(s)
	// [0 1 2 10 10 10]

	n := len(t) // need to get length because t keeps growing
	for i := 0; i < n; i++ {
		t = append(t, 10)
	}
	fmt.Println(t)
	// [0 1 2 10 10 10]

	// Channels 
	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 10
		ch2 <- 11
		ch2 <- 12
		// close(ch2)
	}()
	
	ch := ch1 
	// v gets copied to ch1 each time
	for v := range ch { // no i for channels 
		fmt.Println(v)
		// 0
		// 1
		// 2
		ch = ch2
	}
	close(ch) // closes second channel
	
	// Array 
	a := [3]int{1, 2, 3}
	for i, v := range a {
		a[2] = 10 // updating the original array, not the copy 
		if i == 2 {
			fmt.Println(v) // this is the copy of the 3 position of the array
			// 3
		}
	}
	
	// solution loop through a pointer or print the original array 
	// for i := range a {
	// 	a[2] = 10 
	// 	if i == 2 {
	// 		fmt.Println(a[2])
	// 		// 10 
	// 	}
	// }
	
	// for i, v := range &a { // never copies the full array, might be more efficient
	// 	a[2] = 10 
	// 	if i == 2 {
	// 		fmt.Println(v)
	// 		// 10 
	// 	}
	// }
}
