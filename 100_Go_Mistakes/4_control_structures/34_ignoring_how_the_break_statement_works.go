package main

import "fmt"

func main() {
	// does not break when reaching 2
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		// 0 1 2 3 4
		switch i {
		default:
		case 2:
			break // terminates the execution of the inermost for, switch, or select 
		}
	}
	fmt.Println()	
	// solution, applies with select  
	// same idea with continue 
	loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		// 0 1 2 
		switch i {
			default:
			case 2:
			break loop
		}
	}
}
