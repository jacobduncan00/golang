package main

import (
	"fmt"
	"math"
)

// Vertex is a struct
type Vertex struct {
	X, Y float64
}

// Abs returns float64
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale returns nothing
func (v *Vertex) Scale(f float64) {
	// Have to do it like this cuz it does have *= like C++
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
