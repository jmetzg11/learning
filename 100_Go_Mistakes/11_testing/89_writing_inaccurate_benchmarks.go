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


// Not being careful about compiler optimizations
// popcnt has no side effects. If its result is thrown away, the compiler can
// inline it and delete the call entirely (dead-code elimination) — the loop
// benchmarks nothing, giving a bogus ~0 ns/op.
// Fix: store the result in a package-level var so the compiler can't prove
// it's unused and must keep the call.
//   BenchmarkPopcnt1   ~0.3 ns/op   <- optimized away, lying
//   BenchmarkPopcnt2   ~1   ns/op   <- real cost
const (
	m1  = 0x5555555555555555
	m2  = 0x3333333333333333
	m4  = 0x0f0f0f0f0f0f0f0f
	h01 = 0x0101010101010101
)

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

// Bad: result discarded -> compiler can erase the call
func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}

// good
var global uint64 // package-level: compiler can't tell it won't be read elsewhere

func BenchmarkPopcnt2(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = popcnt(uint64(i)) // keep latest result...
	}
	global = v // ...and publish it so the loop body can't be dropped
}


// Being fooled by the observer effect
// 512-wide rows are exactly 4KB apart in memory → all rows compete for the same
// cache slots → data keeps getting evicted → doesn't fit in cache → slow.
// 513-wide rows are 4104 bytes apart → no collision → stays in cache → fast.
// Bad benchmarks reuse the same matrix every iteration, amplifying this effect.
// Fix: recreate the matrix each iteration so both benchmarks start cold.
func calculateSum512(s [][512]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}

func calculateSum513(s [][513]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}


// want to test which one is faster 
const rows = 1000

var res int64

func createMatrix512(r int) [][512]int64 {
	return make([][512]int64, r)
}

func createMatrix513(r int) [][513]int64 {
	return make([][513]int64, r)
}
// Bad: same matrix reused every iteration — cache is warm after first pass, hides real cost
func BenchmarkCalculateSum512_1(b *testing.B) {
	var sum int64
	s := createMatrix512(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum512(s)
	}
	res = sum
}

func BenchmarkCalculateSum513_1(b *testing.B) {
	var sum int64
	s := createMatrix513(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum513(s)
	}
	res = sum
}

// Good: recreate matrix each iteration so cache starts cold every time
func BenchmarkCalculateSum512_2(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix512(rows)
		b.StartTimer()
		sum = calculateSum512(s)
	}
	res = sum
}

func BenchmarkCalculateSum513_2(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix512(rows)
		b.StartTimer()
		sum = calculateSum512(s)
	}
	res = sum
}

