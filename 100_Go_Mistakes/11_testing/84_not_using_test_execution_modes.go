// The Parallel flag 
// mark that a test has to be run in parallel 
func TestFoo(t *testing.T) {
	t.Parallel()
	// ...
}

// all sequential tests are ran first 
// maximum number of tests that can run simultaneously equals the GOMAXPROCS value

// can set the max number of parallel test with:
go test -parallel 16 .


// The -shuffle flag
// used to randomize the tests 

// if a test failed in whith shuffle we can repeat the random order with the seed value 
go test -shuffle=1636399552801504000 -v . 