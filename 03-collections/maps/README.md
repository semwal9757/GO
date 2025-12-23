# Maps in Go

This example demonstrates maps in Go, which are unordered collections of key-value pairs (similar to dictionaries or hash tables in other languages).

## Concepts Covered

### Declaring and Initializing Maps
Maps use the `map[KeyType]ValueType` syntax:

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Eve":   22,
}
```

### Accessing Elements
Access values using their keys:

```go
fmt.Println("Alice's age:", ages["Alice"])
```

If a key doesn't exist, you get the zero value for the value type (0 for int, "" for string, etc.).

### Adding and Modifying Elements
Simply assign to a key:

```go
ages["Alice"] = 26      // Update existing
ages["Charlie"] = 28    // Add new
```

### Deleting Elements
Use the built-in `delete()` function:

```go
delete(ages, "Eve")
```

### Iterating Over Maps
Use `range` to iterate over key-value pairs:

```go
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

**Note:** Map iteration order is **not guaranteed** and may vary between runs.

### Checking Key Existence
Use the two-value assignment to check if a key exists:

```go
if age, ok := ages["Eve"]; ok {
    fmt.Println("Eve exists:", age)
} else {
    fmt.Println("Eve not found")
}
```

- `age` receives the value (or zero value if key doesn't exist)
- `ok` is a boolean indicating whether the key exists

## Running the Code

```bash
go run maps.go
```

## Expected Output

```
Ages map: map[Alice:25 Bob:30 Eve:22]
Alice's age: 25
Updated map: map[Alice:26 Bob:30 Charlie:28 Eve:22]
After deletion: map[Alice:26 Bob:30 Charlie:28]
Alice is 26 years old
Bob is 30 years old
Charlie is 28 years old
Eve not found
```

**Note:** The order of elements in the printed map may vary.

## Key Points

- **Reference Type:** Maps are references; assigning a map to another variable doesn't copy the data
- **Zero Value:** The zero value of a map is `nil`. A `nil` map cannot store values
- **Unordered:** Maps do not maintain insertion order
- **Dynamic:** Maps grow automatically as you add elements
- **Key Requirements:** Keys must be comparable types (==, !=)
- **Thread Safety:** Maps are not safe for concurrent use without synchronization

## Common Map Operations

```go
// Create with make
m := make(map[string]int)

// Check if map is nil
if m == nil {
    // map is nil
}

// Get length
length := len(m)

// Clear all elements (Go 1.21+)
clear(m)

// Check and get in one line
age := ages["Bob"]  // Returns 0 if "Bob" doesn't exist

// Safe check before use
if age, exists := ages["Bob"]; exists {
    // Use age
}
```

## When to Use Maps

Maps are perfect for:
- Looking up values by key
- Counting occurrences
- Implementing caches
- Storing configuration
- Grouping data

## Performance Notes

- Average O(1) time complexity for lookups, inserts, and deletes
- Maps automatically grow but don't shrink automatically
- Pre-allocate with `make(map[K]V, capacity)` if you know the size