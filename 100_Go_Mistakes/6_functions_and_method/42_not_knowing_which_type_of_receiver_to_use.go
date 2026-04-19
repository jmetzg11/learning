package main

import "fmt"

func main() {
	c := customer{balance: 100.}
	c.addValue(50.)
	fmt.Printf("blanace: %.f\n", c.balance)
	// 100 -> remains unchanged 
	
	c.addPointer(50.)
	fmt.Printf("blanace: %.f\n", c.balance)
	// 150 -> changes the original customer
	
	c1 := customerD{data: &data{
		balance: 100,
	}}
	c1.add(50.)
	fmt.Printf("blanace: %.f\n", c1.data.balance)
	// 150
}

type customer struct {
	balance float64
}

func (c customer) addValue(v float64) { // value receiver
	c.balance += v
}

func (c *customer) addPointer(v float64) { // pointer receiver
	c.balance += v
}

// Receiver must be a pointer
type slice []int 

func (s *slice) add(element int) {
	*s = append(*s, element)
}

// Receiver should be a poitner when the receiver is a large object 

// Receiver must be a value if we have to enforce a receiver's immutability or if the receiver is a map, funcion, or channle 

// Receiver should be a value if the slice doens't have to be mutated, naturally a value type without mutable fields,
// reciever is basyc type such as int, float64 or string 


// Example: don't need a pointer to mutate balance because the fields are nested in a struct 
type customerD struct {
	data *data
}

type data struct {
	balance float64
}

func (c customerD) add(operation float64) {
	c.data.balance += operation
}