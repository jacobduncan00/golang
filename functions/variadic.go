package main

import "fmt"

func sum(nums ...int) {
	// Print the slice and a space after
	fmt.Print(nums, " ")
	// Init total to 0
	total := 0
	for _, num := range nums {
		// iterate through slice and add to total
		total += num
	}
	// Print the product of all values in slice
	fmt.Println(total)
}

func main() {
	sum(1, 2)    // 3
	sum(1, 2, 3) // 6

	// Create a new slice of integers with values listed bollow
	nums := []int{1, 2, 3, 4}
	// Sum the slice by "spreading" the values in the slice into the args for sum function
	sum(nums...) // 10
}
