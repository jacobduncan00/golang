package main

import "fmt"

func main() {
	a := []string{"A", "B", "C"}
	s := a[len(a)-1]
	fmt.Println(a)
	fmt.Printf("%s is the last element in the slice above\n", s)

	// Removing the last element
	a = a[:len(a)-1]
	fmt.Println(a)
}
