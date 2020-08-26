package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Basic input / output in Go
func main() {
	// := lets go lang imply the type
	scanner := bufio.NewScanner(os.Stdin)

	// Whatever we type in as input will automatically be a string
	fmt.Printf("What year were you born? ")
	scanner.Scan()

	// Passing the text we want to parse to an int, we need to pass a base, and then the size of integer
	// strconv returns 2 values, the second one is an err
	// How do we convert year to a number?
	year, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	result := 2020 - year
	fmt.Printf("You will be %d years old at the end of 2020", result)
}
