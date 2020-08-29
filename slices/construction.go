package main

import "fmt"

func main() {
	var s []int                  // creates a nil slice
	s1 := []string{"foo", "bar"} // creates a slice of strings
	s2 := make([]int, 2)         // same as []int{0, 0}
	// make takes a length and an optional capacity
	// 2 is length of slice and 4 is capacity
	s3 := make([]int, 2, 4) // same as new([4]int)[:2]
	fmt.Println(s, s1, s2, s3)
	// len() gets length and cap() gets capacity
	fmt.Println(len(s3), cap(s3))
}
