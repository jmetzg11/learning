package main

import "fmt"

func problem() {
	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			messageCh <- i
		}
		disconnectCh <- struct{}{}
	}()

	for {
		// select has no order an random selection from available communications
		// this prevents a fast sender communication from only being used
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func solutionOne() {
	// unbuffered channel
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			messageCh <- i
		}
		disconnectCh <- struct{}{}
	}()

	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func solutionTwo() {
	messageCh := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			messageCh <- i
		}
		close(messageCh) // disconnt channel replaces with this logic
	}()

	for v := range messageCh {
		fmt.Println(v)
	}
	fmt.Println("disconnection, return ")
}

func solutionThree() {
	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			messageCh <- i
		}
		disconnectCh <- struct{}{}
	}()

	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh: // prioritize the first case over default
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}

func main() {
	fmt.Println("--- Problem (may skip messages) ---")
	problem()
	fmt.Println("--- Solution One ---")
	solutionOne()
	fmt.Println("--- Solution Two ---")
	solutionTwo()
	fmt.Println("--- Solution Three ---")
	solutionThree()
}
