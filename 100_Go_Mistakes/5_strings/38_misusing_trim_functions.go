package main

import (
	"fmt"
	"strings"
)

func main() {
	// removes all runes contained in a give set (e.g. "x" and "o")
	fmt.Println(strings.TrimRight("123oxo", "xo"))
	// 123
	
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))
	// 123o
	
	// similar story with TrimLeft and TrimPrefix 
	
	// strings.Trime is like TrimRight and TrimLeft combined
	fmt.Println(strings.Trim("oxox123oxox", "xo"))
	/// 123
}
