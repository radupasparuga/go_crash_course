package main

import "fmt"

func main() {
	name, email := "Radu", "radupasparuga25@gmail.com" // Shorthand
	var age int32 = 37
	var isCool = true
	isCool = false
	size := 1.8
	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", size)
}
