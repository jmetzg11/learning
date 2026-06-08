// Benchmarks measure how fast code runs
// run with: go test -bench=.
// output:   BenchmarkFoo-8    1000000    1234 ns/op
//                             (runs)     (time per run)
func BenchmarkFoo(b *testing.B) {
	// b.N is chosen by the framework and increased until results stabilize — don't pick it yourself
	for i := 0; i < b.N; i++ {
		foo()
	}
}


// Not resetting the timer 
func BenchmarkFoo(b *testing.B) {
	expensiveSetup()
	b.REsetTimer() // reset timer 
	for i := 0; i < b.N; i++ {
		functionUnderTest()
	}
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		expensiveSetup()
		b.starTimer()
		functionUnderTest()
	}
}


// Make wrong assumption about micro-benchmarks
// many factors may affect performacne 
// use benchstat
func BenchmarkAtomicStoreInt32(b *testing.B) {
	var v int32
	for i := 0; i < b.N; i++ {
		atomic.StoreInt32(&v, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var v int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&v, 1)
	}
}


// Not being carefule about compiler optimizations 
