package main

import "fmt"

func main() {
	// Declare a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Eve":   22,
	}
	fmt.Println("Ages map:", ages)

	// Accessing elements
	fmt.Println("Alice's age:", ages["Alice"])

	// Modify or add elements
	ages["Alice"] = 26
	ages["Charlie"] = 28
	fmt.Println("Updated map:", ages)

	// Delete an element
	delete(ages, "Eve")
	fmt.Println("After deletion:", ages)

	// Iterate over map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Check if key exists
	if age, ok := ages["Eve"]; ok {
		fmt.Println("Eve exists:", age)
	} else {
		fmt.Println("Eve not found")
	}
}
