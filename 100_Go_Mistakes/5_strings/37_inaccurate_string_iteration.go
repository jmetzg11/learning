package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ê is not printed, we jump over position 2, len is 6 
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
		// it prints the UTF-8 representation of the byte at index i
	}
	// position 0: h
	// position 1: Ã
	// position 3: l
	// position 4: l
	// position 5: o
	fmt.Printf("len=%d\n", len(s))
	// len=6
	fmt.Println(utf8.RuneCountInString(s))
	// 5
	
	// Solution - print the value 
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	// position 0: h
	// position 1: ê
	// position 3: l
	// position 4: l
	// position 5: o
	
	// Solution - convert string into a slice of runes 
	runes := []rune(s)
	for i := range runes {
		fmt.Printf("position %d: %c\n", i, runes[i])
	}
	// position 0: h
	// position 1: ê
	// position 2: l
	// position 3: l
	// position 4: o
	
	// print 4th rune
	r := []rune(s)[4]
	fmt.Printf("%c\n", r)
}
