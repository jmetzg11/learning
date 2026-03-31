type Foo struct {
	Bar // Embedded field
}

type Bar struct {
	Baz int
}

foo := Foo{}
foo.Baz = 42

Foo.Baz == Foo.Bar.Baz // Baz is a promoted field

// Wrong usage of Embedded field
package inment

import "sync"

type InMem struct {
	sync.Mutex // Embedded field
	m map[string]int
}

func New() *InMem {
	return &InMem(m: make(map[string]int))
}

func (i *InMem) Get(key string) (int, bool) {
	i.Lock() // Accesses the Lock method directly
	v, contains := i.m[key]
	i.Unlock() // Again, direct access
	return v, contains
}

// caller in different package
m := inmen.New()
m.Lock() // because sync.Mutex is a promoted field it can be accessed by the caller

// Fix
type InMem struct {
	mu sync.Mutex
	m map[string]int
}

// When to use Embedded fields
// Bad
type Logger struct {
	writeCloser io.WriteCloser
}

func (l Logger) Write(p []byte) (int, error) {
	return l.writeCloser.Write(p)
}

func (l Logger) Close() err {
	return l.writeCloser.Close()
}

func main() {
	l := Logger{writeCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
// Good
type Logger struct {
	io.WriteCloser // Embedded field
}

func main() {
	l := Logger{WriteCloser: os.Stdout}
	_, _ = l.Write([]byte("foo"))
	_ = l.Close()
}
