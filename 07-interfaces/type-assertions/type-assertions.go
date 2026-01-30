package main

import (
	"fmt"
	"math"
)

// Animal interface for demonstration
type Animal interface {
	Speak() string
	Move() string
}

// Dog implements Animal
type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Running on four legs"
}

// Cat implements Animal
type Cat struct {
	Name  string
	Color string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Sneaking quietly"
}

// Bird implements Animal
type Bird struct {
	Name    string
	Species string
}

func (b Bird) Speak() string {
	return "Chirp!"
}

func (b Bird) Move() string {
	return "Flying through the air"
}

// Shape interface for geometric shapes
type Shape interface {
	Area() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// assertAnimal demonstrates type assertion with Animal interface
func assertAnimal(a Animal) {
	fmt.Printf("\nAnimal speaks: %s\n", a.Speak())
	fmt.Printf("Animal moves: %s\n", a.Move())

	// Type assertion to check if it's a Dog
	if dog, ok := a.(Dog); ok {
		fmt.Printf("✓ This is a dog named %s, breed: %s\n", dog.Name, dog.Breed)
	}

	// Type assertion to check if it's a Cat
	if cat, ok := a.(Cat); ok {
		fmt.Printf("✓ This is a cat named %s, color: %s\n", cat.Name, cat.Color)
	}

	// Type assertion to check if it's a Bird
	if bird, ok := a.(Bird); ok {
		fmt.Printf("✓ This is a bird named %s, species: %s\n", bird.Name, bird.Species)
	}
}

// describeShape uses type assertion to provide detailed information
func describeShape(s Shape) {
	fmt.Printf("\nShape area: %.2f\n", s.Area())

	// Type assertion with safety check
	if rect, ok := s.(Rectangle); ok {
		fmt.Printf("✓ Rectangle: Width=%.2f, Height=%.2f\n", rect.Width, rect.Height)
		fmt.Printf("  Perimeter: %.2f\n", 2*(rect.Width+rect.Height))
	}

	if circle, ok := s.(Circle); ok {
		fmt.Printf("✓ Circle: Radius=%.2f\n", circle.Radius)
		fmt.Printf("  Circumference: %.2f\n", 2*math.Pi*circle.Radius)
	}
}

// processValue demonstrates type assertion with empty interface
func processValue(value interface{}) {
	fmt.Printf("\nProcessing value: %v (type: %T)\n", value, value)

	// Type assertion for different types
	switch v := value.(type) {
	case int:
		fmt.Printf("✓ Integer: %d, Doubled: %d\n", v, v*2)
	case string:
		fmt.Printf("✓ String: '%s', Length: %d\n", v, len(v))
	case float64:
		fmt.Printf("✓ Float: %.2f, Squared: %.2f\n", v, v*v)
	case bool:
		fmt.Printf("✓ Boolean: %t, Negated: %t\n", v, !v)
	case []int:
		fmt.Printf("✓ Integer slice: %v, Length: %d\n", v, len(v))
	case Dog:
		fmt.Printf("✓ Dog: %s says %s\n", v.Name, v.Speak())
	case *Dog:
		fmt.Printf("✓ Dog pointer: %s says %s\n", v.Name, v.Speak())
	default:
		fmt.Printf("✗ Unknown type: %T\n", v)
	}
}

// safeTypeAssertion demonstrates safe type assertion
func safeTypeAssertion(value interface{}) {
	// Safe assertion with ok pattern
	if str, ok := value.(string); ok {
		fmt.Printf("Successfully converted to string: %s\n", str)
	} else {
		fmt.Printf("Value is not a string, it's a %T\n", value)
	}
}

// unsafeTypeAssertion demonstrates unsafe type assertion (can panic)
func unsafeTypeAssertion(value interface{}) {
	// This will panic if value is not a string
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("⚠️  Panic recovered: %v\n", r)
		}
	}()

	str := value.(string) // Unsafe - no ok check
	fmt.Printf("String value: %s\n", str)
}

// getAnimalInfo uses type switch for cleaner code
func getAnimalInfo(a Animal) string {
	switch animal := a.(type) {
	case Dog:
		return fmt.Sprintf("Dog: %s (Breed: %s)", animal.Name, animal.Breed)
	case Cat:
		return fmt.Sprintf("Cat: %s (Color: %s)", animal.Name, animal.Color)
	case Bird:
		return fmt.Sprintf("Bird: %s (Species: %s)", animal.Name, animal.Species)
	default:
		return "Unknown animal"
	}
}

// complexTypeAssertion demonstrates nested type assertions
func complexTypeAssertion(value interface{}) {
	fmt.Printf("\n--- Complex Type Assertion ---\n")

	// First, check if it's a map
	if m, ok := value.(map[string]interface{}); ok {
		fmt.Println("✓ Value is a map:")
		for key, val := range m {
			// Assert types of values in the map
			switch v := val.(type) {
			case string:
				fmt.Printf("  %s: '%s' (string)\n", key, v)
			case int:
				fmt.Printf("  %s: %d (int)\n", key, v)
			case float64:
				fmt.Printf("  %s: %.2f (float64)\n", key, v)
			case bool:
				fmt.Printf("  %s: %t (bool)\n", key, v)
			default:
				fmt.Printf("  %s: %v (%T)\n", key, v, v)
			}
		}
	} else {
		fmt.Printf("✗ Value is not a map, it's a %T\n", value)
	}
}

func main() {
	fmt.Println("=== Type Assertions in Go ===")

	// 1. Basic type assertions with Animal interface
	fmt.Println("\n--- Animal Type Assertions ---")
	dog := Dog{Name: "Buddy", Breed: "Golden Retriever"}
	cat := Cat{Name: "Whiskers", Color: "Orange"}
	bird := Bird{Name: "Tweety", Species: "Canary"}

	assertAnimal(dog)
	assertAnimal(cat)
	assertAnimal(bird)

	// 2. Type assertions with Shape interface
	fmt.Println("\n--- Shape Type Assertions ---")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	describeShape(rect)
	describeShape(circle)

	// 3. Type switch with empty interface
	fmt.Println("\n--- Type Switch Examples ---")
	processValue(42)
	processValue("Hello, Go!")
	processValue(3.14)
	processValue(true)
	processValue([]int{1, 2, 3, 4, 5})
	processValue(dog)
	processValue(&dog)

	// 4. Safe vs Unsafe type assertions
	fmt.Println("\n--- Safe Type Assertion ---")
	safeTypeAssertion("This is a string")
	safeTypeAssertion(123)

	fmt.Println("\n--- Unsafe Type Assertion ---")
	unsafeTypeAssertion("This is a string")
	unsafeTypeAssertion(123) // This will panic and be recovered

	// 5. Type switch for cleaner code
	fmt.Println("\n--- Type Switch for Animal Info ---")
	animals := []Animal{dog, cat, bird}
	for _, animal := range animals {
		info := getAnimalInfo(animal)
		fmt.Printf("%s - %s\n", info, animal.Speak())
	}

	// 6. Complex nested type assertions
	fmt.Println("\n--- Nested Type Assertions ---")
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"salary":  75000.50,
		"active":  true,
		"hobbies": []string{"reading", "coding"},
	}
	complexTypeAssertion(data)

	// 7. Type assertion with nil interface
	fmt.Println("\n--- Nil Interface Assertion ---")
	var nilInterface interface{}
	if nilInterface == nil {
		fmt.Println("✓ Interface is nil")
	}

	// Attempting to assert on nil interface
	if _, ok := nilInterface.(string); !ok {
		fmt.Println("✓ Cannot assert type on nil interface")
	}

	// 8. Checking interface implementation
	fmt.Println("\n--- Interface Implementation Check ---")
	var a Animal
	a = dog

	// Check if the value implements Animal interface
	if _, ok := a.(Animal); ok {
		fmt.Println("✓ Value implements Animal interface")
	}

	// 9. Type assertion in function returns
	fmt.Println("\n--- Type Assertion in Returns ---")
	getValue := func() interface{} {
		return 42
	}

	if num, ok := getValue().(int); ok {
		fmt.Printf("✓ Got integer: %d\n", num)
	}

	// 10. Practical example: Configuration parser
	fmt.Println("\n--- Configuration Parser ---")
	config := map[string]interface{}{
		"host":    "localhost",
		"port":    8080,
		"timeout": 30.5,
		"debug":   true,
	}

	// Extract and use typed values
	if host, ok := config["host"].(string); ok {
		fmt.Printf("Server host: %s\n", host)
	}

	if port, ok := config["port"].(int); ok {
		fmt.Printf("Server port: %d\n", port)
	}

	if timeout, ok := config["timeout"].(float64); ok {
		fmt.Printf("Timeout: %.1f seconds\n", timeout)
	}

	if debug, ok := config["debug"].(bool); ok {
		fmt.Printf("Debug mode: %t\n", debug)
	}

	fmt.Println("\n✅ Type assertions allow you to extract concrete types from interfaces")
	fmt.Println("⚠️  Always use the 'comma ok' idiom for safe type assertions")
}
