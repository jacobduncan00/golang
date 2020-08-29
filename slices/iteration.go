package main

import "fmt"

func main() {
	s := []string{"Foo", "Bar"}
	// The index OR value iteration value is optional
	// One must be present to iterate but both are not needed
	for index, value := range s {
		fmt.Println(index, value)
	}
}
