package main

import "fmt"

func main() {
	// make method to "make" a string slice with 3 empty value spots and assign to s
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// Format is weird for this
	// set s to append method with arguments of self and things to append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Make another slice with the empty values the length of the s slice
	c := make([]string, len(s))
	// copy s slice values to c
	copy(c, s)
	// so now c has the same values in the slice as s
	fmt.Println("cpy:", c)

	// Slicing a slice LUL
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// What does this do?
	// This creates a slice with values "g", "h", and "i"
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
