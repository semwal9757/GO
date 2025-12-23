# Arrays in Go

This example demonstrates arrays in Go, which are fixed-size sequences of elements of the same type.

## Concepts Covered

### Declaring and Initializing Arrays
Arrays in Go have a fixed size that's part of their type:

```go
var numbers [5]int = [5]int{10, 20, 30, 40, 50}
```

The size `[5]` is part of the array's type, meaning `[5]int` and `[10]int` are different types.

### Accessing Elements
Array elements are accessed using zero-based indexing:

```go
fmt.Println("First element:", numbers[0])
fmt.Println("Last element:", numbers[4])
```

### Modifying Elements
Arrays are mutable; you can change individual elements:

```go
numbers[2] = 35
```

### Iterating with Range
The `range` keyword provides a clean way to iterate over arrays:

```go
for i, value := range numbers {
    fmt.Printf("Index %d => %d\n", i, value)
}
```

- `i` is the index
- `value` is the element at that index
- Use `_` to ignore either value: `for i, _ := range numbers` or `for _, value := range numbers`

## Running the Code

```bash
go run arrays.go
```

## Expected Output

```
Array: [10 20 30 40 50]
First element: 10
Last element: 50
Updated Array: [10 20 35 40 50]
Index 0 => 10
Index 1 => 20
Index 2 => 35
Index 3 => 40
Index 4 => 50
```

## Key Points

- **Fixed Size:** Array size is determined at declaration and cannot change
- **Value Type:** Arrays are values, not references. Assigning one array to another copies all elements
- **Type System:** `[5]int` and `[10]int` are completely different types
- **Zero Values:** Uninitialized array elements have their zero value (0 for int, "" for string, etc.)
- **Length Function:** Use `len(array)` to get the array length

## When to Use Arrays

Arrays are less common in Go than slices. Use arrays when:
- You know the exact size at compile time
- You need the performance of stack allocation
- You want value semantics (copying behavior)

For most cases, **slices** are more flexible and idiomatic in Go.