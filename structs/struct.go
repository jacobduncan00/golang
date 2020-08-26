package main

import "fmt"

// This person struct type has name and age fields
type person struct {
	name string
	age  int
}

// This newPerson function constructs a new person structu with the given name
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// Can return a pointer to a local variable as it will survive the scope of the function
	return &p
}

func main() {
	// This syntax creates a new struct
	fmt.Println(person{"Bob", 20})
	// You can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})
	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})
	// A & prefix will yield a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})
	// Dumb way to do the above thing
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// Struct fields can be accessed with dot notation
	fmt.Println(s.name)

	sp := &s
	// You can also use dot notation with struct pointers as they are automatically dereferenced
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
