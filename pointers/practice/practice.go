package main

import "fmt"

// Get the argument as a pointer to an integer
// Going to be returning an integer
func manipulate(arg *int) int {
	// Create a new variable and assign it to the value of the pointer
	newVal := *arg
	// Return the variable with 99 added to it
	return newVal + 99
}

// Get the argument as a pointer to an integer
func changeAddressValue(arg *int) {
	// Set the pointer value to 0, changing not only the function scope but at the memory address
	*arg = 0
}

func main() {
	i := 9

	i = manipulate(&i)
	fmt.Println(i)

	j := 10
	changeAddressValue(&j)
	fmt.Println(j)
}
