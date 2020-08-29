package main

import (
	"fmt"
)

// return index of value if in slice
// else return length of the slice
func Find(a []string, x string) int {
	for index, value := range a {
		if x == value {
			return index
		}
	}
	return len(a)
}

func BinaryFind(a []string, x string) bool {
	low := 0
	high := len(a) - 1

	for low <= high {
		median := (low + high) / 2

		if a[median] < x {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(a) || a[low] != x {
		return false
	}

	return true

}

// return true if value is in slice
// return false if it is not
func Contains(a []string, x string) bool {
	for _, value := range a {
		if x == value {
			return true
		}
	}
	return false
}

func main() {

	a := []string{"Hello", "World", "Jacob", "Go"}
	searchValue := "Go"

	result := Contains(a, searchValue)
	fmt.Println(result)

	result2 := Find(a, searchValue)
	fmt.Printf("%s found at index %d which is %s\n", searchValue, result2, a[result2])

	searchValue2 := "Hey"
	result3 := BinaryFind(a, searchValue2)
	fmt.Printf("It is %t that %s was found in the slice\n", result3, searchValue2)

	result4 := BinaryFind(a, "World")
	fmt.Printf("It is %t that %s was found in the slice\n", result4, "World")
}
