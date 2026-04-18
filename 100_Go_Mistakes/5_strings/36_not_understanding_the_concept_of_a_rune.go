package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(len(s))
	// 5

	s = "汉"
	fmt.Println(len(s))
	// 3

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Printf("%s\n", s)
	// 汉
}
