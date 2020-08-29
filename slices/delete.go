package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	a[i] = a[len(a)-1] // a[2] = "E" copying last index to index i
	a[len(a)-1] = ""   // Erase last element
	a = a[:len(a)-1]   // Truncate slice

	fmt.Println(a)

	// Slow version

	b := []string{"A", "B", "C", "D", "E"}
	r := 2

	// Remove the element at index i from a.
	copy(b[r:], b[r+1:]) // Shift a[i+1:] left one index
	b[len(b)-1] = ""     // Erase last element
	b = b[:len(b)-1]     // Truncate slice.

	fmt.Println(b)
}
