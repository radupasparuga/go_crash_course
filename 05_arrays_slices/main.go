package main

import (
	"fmt"
)

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"

	// Declare and assign
	arr := [2]string{"carrot", "orange"}

	// Slices
	slice := []string{"uwu", "owo"}
	fmt.Println(fruitArr)
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(fruitArr[0], fruitArr[1])
}
