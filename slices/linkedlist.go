package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	queue.PushBack("Hello ") // Enqueue
	queue.PushBack("world!")

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Print(e.Value) // Print the value at the element
		queue.Remove(e)    // Dequeue
	}
}
