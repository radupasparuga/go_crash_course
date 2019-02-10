package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 1
	if x < y {
		fmt.Printf("%d is less than %d", x, y)
	} else if x > y {
		fmt.Printf("%d is higher than %d", y, x)
	} else if y == x {
		fmt.Printf("%d is equal to %d \n", x, y)
	}

	// Switch
	color := "cien"
	switch color {
	case "green":
		fmt.Printf("Color is green")
	case "red":
		fmt.Printf("Color is red")
	case "blue":
		fmt.Printf("Color is blue")
	default:
		fmt.Printf("Ya got gnomed")
	}
}
