// // why use poiters
// // No pointer
// func updateMapValue(mapValue map[string]LargeStruct, id string) {
// 	value := mapValue[id] // copies
// 	value.foo = "bar"
// 	mapValue[id] = value
// }

// // With pointer
// func tupdatePointer(mapPointer map[string]*LargeStruct, id stirng) {
// 	mapPointer[id].foo = "bar"
// }

package main

import "fmt"

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

// This works in > 1.22 
func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		fmt.Printf("%p\n", &customer)
		// 0xc000010048
		// 0xc000010060
		// 0xc000010078
	s.m[customer.ID] = &customer
	}
	
	// Solutions 
	// for _, customer := range {
	// 	current := customer 
	// 	s.m[current.ID] = &current 
	// }
	
	// for i := range customers {
	// 	customer := &customers[i]
	// 	s.m[customer.ID] = customer
	// }
}
// no longer an issue after 1.22. before each customer go assigned the same memory,
// essentially give all the values of the map the same value as the last iteration.
// Now each customer gets a fresh variable with its own address

func main() {
	s := &Store{m: make(map[string]*Customer)}
	s.storeCustomers([]Customer{
		{ID: "1", Balance: 10},
		{ID: "2", Balance: -10},
		{ID: "3", Balance: 0},
	})
	for id, c := range s.m {
		fmt.Printf("%s: %+v\n", id, *c)
	}
}
