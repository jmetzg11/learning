// Channel can be unbuffered or beffered

// unbuffered channel - without capacity  
// also called synchronous channel
// sender is blocked until receiver receives data from the channel 
ch1 := make(chan int)
ch2 := make(chan int, 0)

// buffered channel - with capacity 
// can send messages until capacity is full
// does not provide strong synchronization
ch3 := make(chan int, 1)
ch3 <-1 // Non-blocking 
ch3 <-2 // blocking 

// Generally use 1 and the buffered size. can pick other capacity sizes if you want to:
// - tie the channel size to the number of goroutines created
// - if you have rate limit issues




