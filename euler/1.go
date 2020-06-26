package main

import "fmt"

// Euler Problem 1
// ---------------
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	var arr []int
	for x := 0; x < 1000; x++ {
		if (x%3 == 0 || x%5 == 0) && x != 0 {
			arr = append(arr, x)
		}
	}

	var sum int
	for y := 0; y < len(arr); y++ {
		sum += arr[y]
	}

	final_result := fmt.Sprintf("%d is the sum of all multiples of 3 or 5 below 1000", sum)
	fmt.Println(final_result)
}
