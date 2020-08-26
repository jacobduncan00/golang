package main

import "fmt"

type rect struct {
	width, height int
}

// OHHHH, pointer receivers means that it will have the literal syntax for the
// pointer or value type
// these 2 functions below have access to the rect struct because it has the pointer receiver and value receiver for it

// This area method has a receiver type of *rect
// What is a receiver type???
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value receiver types
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go will automatically handle the conversion between values and pointers for method calls
	// You want to use a pointer receiver type to avoid copying on method calls or to allow the method to
	// mutate the receiving struct
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
