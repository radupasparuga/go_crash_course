package main

import "fmt"

func main() {
	// Only method I will ever use
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d \n", i)
	}

	// Fizzbuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
