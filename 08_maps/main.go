package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign key-values
	emails["bob"] = "bob@gmail.com"
	emails["john"] = "john@gmail.com"
	emails["bob1"] = "bob1@gmail.com"

	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	numbers["four"] = 4
	numbers["five"] = 5

	delete(numbers, "one")
	fmt.Println(emails)
	fmt.Println(numbers)
}
