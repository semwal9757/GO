# Type Assertions & Switches

When working with interfaces (especially the empty interface), you often need to retrieve the underlying concrete value.

## 📌 Type Assertion

Access functionality specific to the underlying type that is not part of the interface.

```go
var i interface{} = "hello"

// Unsafe assertion (can panic)
s := i.(string)

// Safe assertion (comma-ok idiom)
s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

## 📌 Type Switches

Handle multiple possible types cleanly.

```go
switch v := i.(type) {
case int:
    fmt.Println("It's an int:", v)
case string:
    fmt.Println("It's a string:", v)
default:
    fmt.Println("Unknown type")
}
```
