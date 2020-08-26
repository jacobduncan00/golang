package main

import "fmt"

func main() {

	// Creating an empty map by using the built in make function
	// Syntax: make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Get a value for a key with name[key] syntax
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// This returns the number of key/value pairs in the map
	fmt.Println("len:", len(m))

	// delete method removes key/value pairs for a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// the optional second return value indicates whether the key was preset in
	// the map
	// we can ignore the first return value with a _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
