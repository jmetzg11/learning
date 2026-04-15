package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for _, a := range accounts { // 'a' is a copy 
		a.balance += 1000 // everything we assign is a copy
		fmt.Println(a)
		// {1100}
		// {1200}
		// {1300}
	}

	fmt.Println(accounts)
	// [{100} {200} {300}]
	
	// solution 1 
	for i := range accounts {
		accounts[i].balance += 1000 
		fmt.Println(accounts)
		// [{1100} {200} {300}]
		// [{1100} {1200} {300}]
		// [{1100} {1200} {1300}]
	}
	
	// solution 2 
	// for i := 0; i < len(accounts); i++ {
	// 	acounts[i].balance += 100
	// }
	
	// solution 3 
	// accounts := []*account{
	// 	{balance: 100.},
	// 	{balance: 200.},
	// 	{balance: 300.},
	// }
	// for _, a := range acounts {
	// 	a.balance += 100
	// }
	// a is a copy of a pointer pointing to the same struct. Not idea because accounts is now a reference and slices of pointers aren't as efficient
}
