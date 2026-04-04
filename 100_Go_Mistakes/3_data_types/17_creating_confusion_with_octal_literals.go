package main

import "fmt"

func main() {
	sum := 100 + 010 // literal starting with 0 is consider an ocatla integer (base 8)
	fmt.Println(sum)
}
// prints 108


// Unix permissions use octal because each digit (0-7) maps to one permission group (owner/group/others).
// Each group has read=4, write=2, execute=1. So 0644 means owner=rw-, group=r--, others=r--.
// The leading 0 (or 0o) tells Go it's octal, otherwise 644 decimal is a completely different number.
file, err := os.OpenFile("foo", os.O_RDONLY, 0644)
file, err := os.OpenFile("foo", os.O_RDONLY, 0o644)


// Binary - 0b / 0B
// Hexadecimal - 0x / 0X
// Imaginary - i (e.g. 3i)


// use underscores when possible; 1_000_000_000
