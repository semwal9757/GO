package main

import "fmt"

func main() {
	// Declare an array of 5 integers
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", numbers)

	// Accessing elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[4])

	// Modify an element
	numbers[2] = 35
	fmt.Println("Updated Array:", numbers)

	// Iterate over array
	for i, value := range numbers {
		fmt.Printf("Index %d => %d\n", i, value)
	}
}
