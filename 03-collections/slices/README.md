# Slices in Go

This example demonstrates slices in Go, which are dynamic, flexible views into arrays. Slices are one of the most important and commonly used data structures in Go.

## Concepts Covered

### Creating Slices
Slices are declared without specifying a size:

```go
fruits := []string{"Apple", "Banana", "Cherry"}
```

Unlike arrays, slices are dynamic and can grow or shrink.

### Appending Elements
Use the `append()` function to add elements to a slice:

```go
fruits = append(fruits, "Date")
```

**Important:** `append()` returns a new slice, so you must assign the result back to the variable.

### Combining Slices
Use the `...` operator to append all elements from one slice to another:

```go
citrus := []string{"Orange", "Lemon", "Lime"}
allFruits := append(fruits, citrus...)
```

### Iterating with Range
Just like arrays, use `range` to iterate over slices:

```go
for i, fruit := range allFruits {
    fmt.Printf("Index %d => %s\n", i, fruit)
}
```

### Length and Capacity
Slices have both length and capacity:

```go
fmt.Println("Length:", len(allFruits))    // Number of elements
fmt.Println("Capacity:", cap(allFruits))  // Size of underlying array
```

- **Length:** Number of elements in the slice
- **Capacity:** Number of elements in the underlying array, starting from the first element in the slice

## Running the Code

```bash
go run slices.go
```

## Expected Output

```
Fruits slice: [Apple Banana Cherry]
After append: [Apple Banana Cherry Date]
Combined slice: [Apple Banana Cherry Date Orange Lemon Lime]
Index 0 => Apple
Index 1 => Banana
Index 2 => Cherry
Index 3 => Date
Index 4 => Orange
Index 5 => Lemon
Index 6 => Lime
Length: 7
Capacity: 8
```

## Key Points

- **Dynamic Size:** Slices can grow and shrink dynamically
- **Reference Type:** Slices are references to underlying arrays; copying a slice doesn't copy the data
- **Zero Value:** The zero value of a slice is `nil`
- **Make Function:** Create slices with specific length/capacity: `make([]int, length, capacity)`
- **Slicing Operation:** Create sub-slices: `slice[start:end]`

## Common Slice Operations

```go
// Create with make
s := make([]int, 5)      // length 5, capacity 5
s := make([]int, 5, 10)  // length 5, capacity 10

// Sub-slicing
s[1:4]   // elements from index 1 to 3
s[:3]    // first 3 elements
s[2:]    // from index 2 to end

// Check for nil
if s == nil {
    // slice is nil
}
```

## Why Slices Are Better Than Arrays

- More flexible and idiomatic in Go
- Can grow dynamically
- Easier to pass to functions (no size in type)
- Most Go standard library functions use slices