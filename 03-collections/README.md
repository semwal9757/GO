# Collections in Go

This section covers Go's primary collection types: arrays, slices, and maps. These data structures are fundamental to writing effective Go programs.

## 📂 Directory Structure

```
03-collections/
├── arrays/
│   ├── arrays.go
│   └── README.md
├── slices/
│   ├── slices.go
│   └── README.md
├── maps/
│   ├── maps.go
│   └── README.md
└── README.md
```

## 📚 Topics Covered

### 1. Arrays (`arrays/`)
Fixed-size sequences of elements:
- Declaration and initialization
- Accessing and modifying elements
- Iterating with `range`
- Value semantics and copying
- When to use arrays vs slices

**Key Feature:** Arrays have a fixed size that's part of their type, making them less flexible but more predictable.

### 2. Slices (`slices/`)
Dynamic, flexible views into arrays:
- Creating and initializing slices
- Appending and combining elements
- Length vs capacity
- Sub-slicing operations
- Reference semantics

**Key Feature:** Slices are the most commonly used collection type in Go, offering dynamic sizing and powerful operations.

### 3. Maps (`maps/`)
Key-value pairs for fast lookups:
- Declaring and initializing maps
- Adding, updating, and deleting entries
- Checking key existence
- Iterating over maps
- Reference semantics and nil maps

**Key Feature:** Maps provide O(1) average-time lookups and are perfect for associative data.

## 🚀 Getting Started

Each subdirectory contains:
- A `.go` file with practical examples
- A `README.md` with detailed explanations

To run any example:

```bash
# Navigate to a specific folder
cd arrays/

# Run the example
go run arrays.go
```

Or run directly from the collections directory:

```bash
go run arrays/arrays.go
go run slices/slices.go
go run maps/maps.go
```

## 🎯 Key Differences

| Feature | Arrays | Slices | Maps |
|---------|--------|--------|------|
| **Size** | Fixed | Dynamic | Dynamic |
| **Type** | `[n]T` | `[]T` | `map[K]V` |
| **Semantics** | Value | Reference | Reference |
| **Common Usage** | Rare | Very Common | Very Common |
| **Zero Value** | Array of zeros | `nil` | `nil` |
| **Growth** | Cannot grow | Grows with `append()` | Grows automatically |

## 📖 Quick Reference

### Arrays
```go
var arr [5]int                    // Declare
arr := [5]int{1, 2, 3, 4, 5}     // Initialize
arr := [...]int{1, 2, 3}          // Size inferred
```

### Slices
```go
s := []int{1, 2, 3}               // Literal
s := make([]int, 5)               // Length 5
s := make([]int, 5, 10)           // Length 5, capacity 10
s = append(s, 4, 5, 6)            // Append elements
```

### Maps
```go
m := map[string]int{}             // Empty map
m := make(map[string]int)         // Using make
m := map[string]int{"a": 1}       // Literal
value, ok := m["key"]             // Check existence
delete(m, "key")                  // Delete entry
```

## 🎓 Best Practices

1. **Prefer Slices Over Arrays:** Unless you have a specific reason, use slices
2. **Initialize Maps:** Always initialize maps before use (nil maps can't store values)
3. **Check Map Keys:** Use the `value, ok` pattern to check key existence
4. **Preallocate When Possible:** Use `make()` with capacity for better performance
5. **Use Range:** The `range` keyword is idiomatic for iteration
6. **Remember Reference Semantics:** Slices and maps are references, not copies

## 💡 Common Patterns

### Filtering a Slice
```go
var filtered []int
for _, v := range numbers {
    if v > 10 {
        filtered = append(filtered, v)
    }
}
```

### Checking Duplicates with a Map
```go
seen := make(map[string]bool)
for _, item := range items {
    if seen[item] {
        fmt.Println("Duplicate:", item)
    }
    seen[item] = true
}
```

### Converting Slice to Set (Unique Values)
```go
set := make(map[string]struct{})
for _, item := range items {
    set[item] = struct{}{}
}
```

## 📘 Learn More

- [Go by Example: Arrays](https://gobyexample.com/arrays)
- [Go by Example: Slices](https://gobyexample.com/slices)
- [Go by Example: Maps](https://gobyexample.com/maps)
- [Effective Go: Slices](https://go.dev/doc/effective_go#slices)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)

## 🔄 What's Next?

After mastering collections, explore:
- Structs and methods
- Interfaces
- Error handling
- Concurrency with goroutines and channels
- Pointers and memory management