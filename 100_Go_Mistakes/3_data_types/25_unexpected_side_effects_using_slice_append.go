package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)
	fmt.Printf("slice 1: %v, slice 3: %v\n", s1, s3)
	// s1 was modified by the append
	// slice 1: [1 2 10], slice 3: [2 10]

	s := []int{1, 2, 3}
	sCopy := make([]int, 2)
	copy(sCopy, s)
	f(sCopy)
	fmt.Printf("original: %v, copy: %v\n", s, sCopy)
}

func f(s []int) {
	_ = append(s, 10)
}
