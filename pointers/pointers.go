// Pointers allow you to pass references to values and records within your program

package main

import "fmt"

// Has an int parameter, so arguments will be passed to it by value
// Will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// Has an *int parameter, meaning that it takes an int pointer.
// The *iptr code in the function body then dereferences the pointer from its memory address
// to the current value at that address. Assigning a value to a dereferenced pointer changes the
// the value at the referenced address
func zeroptr(iptr *int) {
	// Setting the memory address to 0, therefore it is changed not only inside the function, but throught the rest of the program
	*iptr = 0
}

func main() {
	i := 1
	// Obviously prints 1
	fmt.Println("initial:", i)

	zeroval(i)
	// Prints 1 because the value of i does not get changed in the zeroval function
	fmt.Println("zeroval:", i)

	// The &i syntax gives the memory address of i, (a pointer to i)
	zeroptr(&i)
	// Should print 0
	fmt.Println("zeroptr:", i)

	// Pointers can be printed too
	// This will print the memory address of i
	// looks something like 0x...
	fmt.Println("pointer:", &i)
}
