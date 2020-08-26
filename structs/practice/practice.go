package main

import "fmt"

type school struct {
	name       string
	location   string
	population int
}

func newSchool(name string, location string, population int) *school {
	s := school{name: name, location: location, population: population}
	return &s
}

func main() {
	middleSchool := school{"Yokota Middle School", "Tokyo, Japan", 150}
	fmt.Println(middleSchool)
	fmt.Println(newSchool("Bates Middle School", "Sumter, South Carolina", 499))
}
