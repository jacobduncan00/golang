package main

import "fmt"

func main() {
	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ")

	// Print top element, then pop, continue
	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n])

		stack = stack[:n] // Pop
	}
}

// So, whats the performance?
/*
	Appending a single element to a slice takes constant time.
	If the stack is permanent and the elements temporary, you may want to
	remove the top element before popping the stack to avoid memory leaks
*/
