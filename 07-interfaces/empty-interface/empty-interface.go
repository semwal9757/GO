package main

import "fmt"

// The empty interface (interface{}) can hold values of any type
// In Go 1.18+, you can also use 'any' which is an alias for interface{}

// printAnything accepts any type
func printAnything(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// describe provides detailed information about any value
func describe(value any) {
	fmt.Printf("\n--- Describing value ---\n")
	fmt.Printf("Value: %v\n", value)
	fmt.Printf("Type: %T\n", value)
	fmt.Printf("Underlying representation: %#v\n", value)
}

// Container can store any type of value
type Container struct {
	Value interface{}
}

// Set updates the container's value
func (c *Container) Set(value interface{}) {
	c.Value = value
}

// Get retrieves the container's value
func (c *Container) Get() interface{} {
	return c.Value
}

// GenericMap is a map that can store any type of value
type GenericMap map[string]interface{}

// Set adds or updates a key-value pair
func (gm GenericMap) Set(key string, value interface{}) {
	gm[key] = value
}

// Get retrieves a value by key
func (gm GenericMap) Get(key string) (interface{}, bool) {
	value, exists := gm[key]
	return value, exists
}

// processSlice works with a slice of any type
func processSlice(items []interface{}) {
	fmt.Println("\nProcessing slice of mixed types:")
	for i, item := range items {
		fmt.Printf("  [%d] %v (type: %T)\n", i, item, item)
	}
}

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
}

// Config demonstrates using empty interface for configuration
type Config struct {
	Settings map[string]interface{}
}

func (c *Config) Set(key string, value interface{}) {
	if c.Settings == nil {
		c.Settings = make(map[string]interface{})
	}
	c.Settings[key] = value
}

func (c *Config) Get(key string) interface{} {
	return c.Settings[key]
}

// compareValues compares two values of any type
func compareValues(a, b interface{}) {
	fmt.Printf("\nComparing %v (%T) with %v (%T)\n", a, a, b, b)
	if a == b {
		fmt.Println("Values are equal")
	} else {
		fmt.Println("Values are not equal")
	}
}

// JSON-like structure using empty interfaces
func createJSONLike() map[string]interface{} {
	return map[string]interface{}{
		"name":   "John Doe",
		"age":    30,
		"active": true,
		"scores": []interface{}{95, 87, 92},
		"address": map[string]interface{}{
			"street": "123 Main St",
			"city":   "New York",
			"zip":    10001,
		},
	}
}

func main() {
	fmt.Println("=== Empty Interface (interface{} / any) ===")

	// 1. Basic usage
	fmt.Println("\n--- Basic Usage ---")
	printAnything(42)
	printAnything("Hello, Go!")
	printAnything(3.14)
	printAnything(true)
	printAnything([]int{1, 2, 3})

	// 2. Using 'any' (Go 1.18+)
	fmt.Println("\n--- Using 'any' keyword ---")
	var value any
	value = 100
	describe(value)

	value = "String value"
	describe(value)

	value = Person{Name: "Alice", Age: 25}
	describe(value)

	// 3. Container example
	fmt.Println("\n--- Generic Container ---")
	container := Container{}

	container.Set(42)
	fmt.Printf("Container holds: %v (type: %T)\n", container.Get(), container.Get())

	container.Set("Now it's a string")
	fmt.Printf("Container holds: %v (type: %T)\n", container.Get(), container.Get())

	container.Set(Person{Name: "Bob", Age: 30})
	fmt.Printf("Container holds: %v (type: %T)\n", container.Get(), container.Get())

	// 4. Generic map
	fmt.Println("\n--- Generic Map ---")
	gm := make(GenericMap)
	gm.Set("name", "Charlie")
	gm.Set("age", 28)
	gm.Set("salary", 75000.50)
	gm.Set("active", true)

	for key, value := range gm {
		fmt.Printf("%s: %v (type: %T)\n", key, value, value)
	}

	// 5. Slice of different types
	fmt.Println("\n--- Slice of Mixed Types ---")
	mixedSlice := []interface{}{
		42,
		"hello",
		3.14,
		true,
		Person{Name: "Diana", Age: 35},
		[]int{1, 2, 3},
	}
	processSlice(mixedSlice)

	// 6. Configuration example
	fmt.Println("\n--- Configuration System ---")
	config := &Config{}
	config.Set("host", "localhost")
	config.Set("port", 8080)
	config.Set("debug", true)
	config.Set("timeout", 30.5)

	fmt.Println("Configuration:")
	for key, value := range config.Settings {
		fmt.Printf("  %s: %v (type: %T)\n", key, value, value)
	}

	// 7. JSON-like structure
	fmt.Println("\n--- JSON-like Structure ---")
	data := createJSONLike()
	fmt.Println("User data:")
	for key, value := range data {
		fmt.Printf("  %s: %v\n", key, value)
	}

	// Accessing nested values
	if address, ok := data["address"].(map[string]interface{}); ok {
		fmt.Printf("\nCity from nested address: %v\n", address["city"])
	}

	// 8. Comparing values
	fmt.Println("\n--- Comparing Values ---")
	compareValues(42, 42)
	compareValues(42, "42")
	compareValues("hello", "hello")
	compareValues(true, false)

	// 9. Function that returns different types
	fmt.Println("\n--- Dynamic Return Values ---")
	getValue := func(key string) interface{} {
		values := map[string]interface{}{
			"count":   100,
			"message": "Success",
			"ratio":   0.95,
			"enabled": true,
		}
		return values[key]
	}

	fmt.Printf("count: %v (type: %T)\n", getValue("count"), getValue("count"))
	fmt.Printf("message: %v (type: %T)\n", getValue("message"), getValue("message"))
	fmt.Printf("ratio: %v (type: %T)\n", getValue("ratio"), getValue("ratio"))
	fmt.Printf("enabled: %v (type: %T)\n", getValue("enabled"), getValue("enabled"))

	// 10. Variadic function with empty interface
	fmt.Println("\n--- Variadic Function ---")
	printAll := func(items ...interface{}) {
		fmt.Println("Printing all items:")
		for i, item := range items {
			fmt.Printf("  Item %d: %v (type: %T)\n", i+1, item, item)
		}
	}

	printAll(1, "two", 3.0, true, Person{Name: "Eve", Age: 40})

	// Note: While empty interface is powerful, it loses type safety
	// Use it when you truly need to work with values of unknown types
	// Consider using generics (Go 1.18+) for type-safe generic code
	fmt.Println("\n⚠️  Note: Empty interface provides flexibility but loses type safety.")
	fmt.Println("Consider using type assertions or generics for better type safety.")
}
