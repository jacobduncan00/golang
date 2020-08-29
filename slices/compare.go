package main

import "fmt"

// Comparison for slices
func EqualSlice(a, b []int) bool {
	// If lengths aren't the same we already know they aren't the same
	if len(a) != len(b) {
		return false
	}
	// Iterate over a and compare with the value at each index in b
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

func EqualArray(a [2]int, b [2]int) {
	fmt.Println(a == b)
}

func main() {
	a := [2]int{1, 2}
	b := [2]int{1, 3}
	EqualArray(a, b)
	a2 := []int{1, 2}
	b2 := []int{1, 2}
	returnValue := EqualSlice(a2, b2)
	fmt.Println(returnValue)
}
