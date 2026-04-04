package main

// Signed integers:
//   int8    -128 to 127
//   int16   -32,768 to 32,767
//   int32   -2,147,483,648 to 2,147,483,647
//   int64   -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
//
// Unsigned integers:
//   uint8   0 to 255
//   uint16  0 to 65,535
//   uint32  0 to 4,294,967,295
//   uint64  0 to 18,446,744,073,709,551,615

// int and uint have size that depend on the systme: 32 bits or 64 bints

import (
	"fmt"
	"math"
)

func main() {
	var counter int32 = math.MaxInt32
	counter++
	fmt.Printf("counter=%d\n", counter)
	// counter=-2147483648
	// bit overflow, first bit (signed) gets flipped, turning the number negative
}

// detection overflow
func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow")
	}
	return counter + 1
}

func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("int overflow")
	}
	retrun counter + 1
}

func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("int overflow")
	}
	return a + b
}

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}
	if a == math.MinInt || b == math.MinInt { // prevent the division from panicking
		panic("integer overflow")
	}
	if result/b != a { // did int overflow occur
		panic("integer overflow")
	}
	return result
}

