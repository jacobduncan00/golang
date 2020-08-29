package main

import "fmt"

// Need to figure out how the capacity changes here
func main() {
	a := [...]int{0, 1, 2, 3} // an array
	s := a[1:3]               // s == []int{1, 2} cap(s) == 3
	s2 := a[:2]               // s == []int{0, 1} cap(s) == 4
	s3 := a[2:]               // s == []int{2, 3} cap(s) == 2
	s4 := a[:]                // s == []int{0, 1, 2, 3} cap(s) == 4
	fmt.Println(a, s, s2, s3, s4)

	a2 := []int{0, 1, 2, 3, 4} // a slice
	s5 := a2[1:4]              // s == []int{1, 2, 3}
	s6 := a2[1:2]              // s == []int{2} (index relative to slice)
	s7 := a2[:3]               // s == []int{2, 3, 4} (extend length)
	fmt.Println(a2, s5, s6, s7)
}
