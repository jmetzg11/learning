package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := getInput()

	start := time.Now()
	concatBad(s)
	fmt.Printf("concatBad:  %v\n", time.Since(start))
	// 49ms
	start = time.Now()
	concatGood(s)
	fmt.Printf("concatGood: %v\n", time.Since(start))
	// 849us
	start = time.Now()
	concatBest(s)
	fmt.Printf("concatBest: %v\n", time.Since(start))
	// 83us
}

func concatBad(values []string) string {
	s := ""
	for _, value := range values {
		s += value // reallocates a new string in memory, strings are immutabible
	}
	return s
}

// use strings.Builder when concatenation is ~5 or more strings
func concatGood(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

// pre allocating capatcity for slice is always a big win
func concatBest(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}
	sb := strings.Builder{}
	sb.Grow(total) // preallocate space
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func getInput() []string {
	n := 1_000
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = string(make([]byte, 1_000))
	}
	return s
}
