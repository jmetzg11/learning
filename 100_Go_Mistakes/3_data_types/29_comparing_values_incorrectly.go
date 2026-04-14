package main

import (
	"fmt"
	"reflect"
)

type customer struct {
	id string
}

type customer2 struct {
	id         string
	operations []float64
}

func main() {
	cust1 := customer{id: "x"}
	cust2 := customer{id: "x"}
	fmt.Println(cust1 == cust2)
	// true

	cust12 := customer2{id: "x", operations: []float64{1.}}
	cust22 := customer2{id: "x", operations: []float64{1.}}
	// fmt.Println(cust12 == cust22) ->  operation doesn't work with slices or maps
	// invalid operation: cust12 == cust22 (struct containing []float64 cannot be compared)}

	fmt.Println(reflect.DeepEqual(cust12, cust22)) // works with the reflect library, but 100 times slower
	// true

	fmt.Println(cust12.equal(cust22)) // custom function about 96 times faster than reflect.DeepEqual because know the exact type at compile time
	// true

	// careful with any, it will compile but gets an error at run time
}

func (a customer2) equal(b customer2) bool {
	if a.id != b.id {
		return false
	}
	if len(a.operations) != len(b.operations) {
		return false
	}
	for i := 0; i < len(a.operations); i++ {
		if a.operations[i] != b.operations[i] {
			return false
		}
	}
	return true
}
