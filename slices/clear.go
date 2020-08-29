package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a))

	b := []string{"A", "B", "C", "D", "E"}
	b = b[:0]
	fmt.Println(b, len(b), cap(b))

	// In practice, empty and nil slices can often be treated the same way

	var c []int = nil
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println(c)
}
