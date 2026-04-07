package main

import (
	"fmt"
)

func main() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst)
	// []

	dst2 := make([]int, len(src)) // need to have at least lenth of src
	copy(dst2, src)
	fmt.Println(dst2)
	// [0 1 2]

	dst3 := append([]int(nil), src...)
	fmt.Println(dst3)
	// [0 1 2]
}
