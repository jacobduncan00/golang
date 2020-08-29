package main

import "fmt"

func main() {
	var queue []string

	queue = append(queue, "Hello ")
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0]) // Print first element
		queue = queue[1:]   // Dequeue
	}
}

// Possible memory leaks

/*
	You may want to remove the first element before dequeuing.
	queue[0] = "" // Erase element
	queue = queue[1:]
*/
