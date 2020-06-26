// First "Program" in Golang
// 06-25-2020
// Jacob Duncan

package main

import (
	"fmt"
)

func add(x int, y int) int {
	var return_value int = x + y
	return return_value
}

func sub(x int, y int) int {
	var return_sub_value int = x - y
	return return_sub_value
}

func mult(x int, y int) int {
	var return_mult_value int = x * y
	return return_mult_value
}

func div(x int, y int) int {
	var return_div_value int = x / y
	return return_div_value
}

func main() {
	fmt.Println("hi")
	var add_result int = add(10, 10)
	fmt.Println(add_result)
	var mult_result int = mult(9, 7)
	fmt.Println(mult_result)
	var div_result int = div(150, 7)
	fmt.Println(div_result)
	var sub_result int = sub(100, 50)
	fmt.Println(sub_result)
}
