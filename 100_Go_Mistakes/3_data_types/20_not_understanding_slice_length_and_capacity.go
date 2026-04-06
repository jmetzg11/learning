package main

import "fmt"

func main() {
	// length is 3, go only initializes the first 3
	s := make([]int, 3, 6) // memory: [0, 0, 0,,,]
	s[1] = 1
	// event though the slice has the capacity this is forbidden
	// x := s[3]

	// length of slice is now 4
	s = append(s, 2)
	fmt.Println(s)

	// There was no capacity for 5
	// Internally go creates another array
	// and doubles capacity: [0,1,0,2,3,4,5,,,,,]
	// len = 7, cap = 12
	s = append(s, 3, 4, 5)
	fmt.Println(s)

	// slicing
	s1 := make([]int, 3, 6)
	s2 := s1[1:3]
	// change affects both
	s1[1] = 99
	// length of s2 is modified and only s2 is modified
	s2 = append(s2, 2)
	fmt.Printf("slicing: s1: %v, s2: %v\n", s1, s2)

	// s1 and s2 now reference tow different arrays, s1 is unchanged
	// the new array was made by copyting the initial array from the first index of s2
	s2 = append(s2, 3, 4, 5)
	fmt.Printf("s1: %v, s2: %v\n", s1, s2)
}
