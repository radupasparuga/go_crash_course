package main

import "fmt"

func main() {
	ids := []int{33, 25, 21, 5, 11, 10}

	// Looping trough slice

	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// In case I don't need "i"
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Sum of ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with maps
	emails := make(map[string]string)

	// Assign key-values
	emails["bob"] = "bob@gmail.com"
	emails["john"] = "john@gmail.com"
	emails["bob1"] = "bob1@gmail.com"

	for k, v := range emails {
		fmt.Println(k, v)
	}
}
