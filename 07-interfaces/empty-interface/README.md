# The Empty Interface

The empty interface `interface{}` (or `any` in Go 1.18+) has no methods. Since every type implements at least zero methods, **every type implements the empty interface**.

## 📌 Usage

It is often used when you need to handle values of unknown type, similar to `Object` in Java or `any` in TypeScript.

```go
func PrintAny(v interface{}) {
    fmt.Printf("%v\n", v)
}

// Can be called with anything
PrintAny(42)
PrintAny("hello")
PrintAny(MyStruct{})
```

## ⚠️ Traceability & Safety

While powerful, the empty interface opts out of static type checking. Use it sparingly.
- You lose compile-time type safety.
- You often need **Type Assertions** to get the concrete value back.
