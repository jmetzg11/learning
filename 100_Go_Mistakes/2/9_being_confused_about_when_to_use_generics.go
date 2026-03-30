// We don't know m type
func getKeys(m any) ([]any, error) {
	switch t := m.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", t)
	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil
	case map[int]string:
		// ...
	}

	return nil, nil
}

// Generics

// keys can only be comparables while values can be anyting
func getKeysGenerics[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Custom contraings for the key type
type customConstraint interface {
	~int | ~string
}

func getKeysWithConstraint[K customConstraint, V any](m map[K]V) []K {
	return nil
}

// Linked list
type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

// factoring out behaviors instead of types
// some swiss army knife shit
type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }

func main() {
	s := SliceFn[int]{
		S: []int{3, 2, 1},
		Compare: func(a, b int) bool {
			return a < b
		},
	}
	sort.Sort(s)
	fmt.Println(s.S)
}

// Generics are useful
// data structures
// function working with channels
func merge[T any](ch1, ch2 <-chan T) <-chan T {

}

// When not to use generics
// calling a method of the type argument
// when it makes our code more complex
