package main

import "fmt"

func main() {
	// This creates a slice of integers with values 2, 3, 4
	nums := []int{2, 3, 4}
	sum := 0
	// iterate over nums slice with the range as the length
	// assign each index's value to num
	// add that to the sum variable
	for _, num := range nums {
		sum += num
	}
	// print the sum of the slice value, expecting 9
	fmt.Println("sum:", sum)

	// Now we get the index and the value at the index and print just the index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// create a new map of strings and assign key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}

	// grab key value at each index of the map
	// and print each of the corresponding key value pairs

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate and print only the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. The first value
	// is the starting byte index of the rune and the second the rune itself

	// TODO: look into this more
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
