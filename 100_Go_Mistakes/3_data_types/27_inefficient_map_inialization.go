// this map is backed by an array consisting of a single enter, hence a single bucket
m := map[string]int {
	"1": 1, 
	"2": 2, 
	"3": 3,
}

// when a map grows it increases it buckets automatically, doubling in size. Grow determinted by an internal 'load factor'

// if you know the size beforehand you can avoid costly growth operations 

// only able to give initial size, not capacity 
m := make(map[string]int, 1_000_000) // map can still grow, but we at least already allocated enough space to be more efficient 


