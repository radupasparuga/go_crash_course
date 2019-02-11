package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " and I'm " + strconv.Itoa(p.age)
}

func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Init defined struct
	person1 := Person{firstName: "Radu", lastName: "Pasparuga", city: "EEEEOOO", gender: "LGBTQBBY", age: 211}
	// Faster way
	person2 := Person{"bill", "gates", "motherfucking silicon valley", "m", 60}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.greet())
}
