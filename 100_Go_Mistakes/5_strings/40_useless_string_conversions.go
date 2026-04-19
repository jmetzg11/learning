// we prefer working with strings but most I/O is done with []byte

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	b := []byte{'a', 'b', 'c'}
	s := string(b)
	b[1] = 'x'
	fmt.Println(s)
	// abc
}

func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return sanitize(b), nil
}

// would need bytes -> strings -> bytes = not efficient 
func santizieBad(s string) string {
	return strings.TrimSpace(s)
}

func sanitize(b []byte) []byte {
	return bytes.TrimSpace(b) // most string methods work on bytes too
}
