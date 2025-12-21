package main

import "fmt"

func main() {
	// Create a slice
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits slice:", fruits)

	// Append elements
	fruits = append(fruits, "Date")
	fmt.Println("After append:", fruits)

	// Slice from slice
	citrus := []string{"Orange", "Lemon", "Lime"}
	allFruits := append(fruits, citrus...)
	fmt.Println("Combined slice:", allFruits)

	// Iterate over slice
	for i, fruit := range allFruits {
		fmt.Printf("Index %d => %s\n", i, fruit)
	}

	// Length and capacity
	fmt.Println("Length:", len(allFruits))
	fmt.Println("Capacity:", cap(allFruits))
}
