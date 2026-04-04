// floating-point arithmetic is an approximation of real arithmetic

package main

import (
	"fmt"
	"math"
)

func main() {
	var n float32 = 1.0001
	fmt.Println(n * n)
	// outputs 1.0002, instead of 1.00020001

	var a float64
	positiveInf := 1 / a
	negativeInf := -1 / a
	nan := a / a
	fmt.Printf("positiveInf: %v, isInf: %v\n", positiveInf, math.IsInf(positiveInf, 1))
	fmt.Printf("negativeInf: %v, isInf: %v\n", negativeInf, math.IsInf(negativeInf, -1))
	fmt.Printf("nan: %v, isNan: %v\n", nan, math.IsNaN(nan))

	ns := [4]int{10, 1_000, 1_000_000, 1_000_000_000}
	for _, n := range ns {
		f1Result := f1(n)
		f2Result := f2(n)
		fmt.Println(f1Result, f2Result)
	}
}

// there are infinite numbers betwee min float64 and max float64
// but it's impossilbe to make infinite values fit into a finite space
// some bits represent mantissa (base value) and other represent an exponent (multipler applied to mantissa)
// sign * 2^exponent * mantissa

// don't comapre floats, instead compare their difference for a small enough delta

func f1(n int) float64 {
	result := 10_000.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result
}

func f2(n int) float64 { // more accurate because add 10,000 at the end
	result := 0.
	for i := 0; i < n; i++ {
		result += 1.0001
	}
	return result + 10_000.
}

// when performing addition, subtraction, multipliaction, and division - do multipliaction and division to get better accuracy
// but it may affect execution. It's a trade off between performance and accuracy
// group operations of similar magnitude together
