

s := []int{1,2,3}

for _, i := range s {
	go runc() {
		fmt.Print(i)
	}()
}
// will print 233, 113, 333 -> no predictable 
// calling a goroutine in a closure (referencing a variable from outside)
// the go routine access the variable when fmt.Print() is executed 

// solution #1
s := []int{1,2,3}

for _, i := range s {
	val := i 
	go func() {
		fmt.Print(val)
	}()
}

// each closure call has access to the correct local variable 

// solution #2

s := []int{1,2,3}

for _, i := range s {
	go func(val int) {
		fmt.Print(val)
	}(i)
}

// val is now part of the function 