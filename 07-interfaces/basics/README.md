# Interface Basics

Interfaces in Go provide a way to specify the behavior of an object: if something can do *this*, then it can be used *here*.

## 📌 Implementation

In Go, interfaces are implemented **implicitly**. There is no `implements` keyword. If a type provides definitions for all the methods declared in an interface, it is said to implement that interface.

```go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

// Rectangle implicitly implements Shape because it has an Area method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

## 🔑 Key Concepts

- **Definition**: A collection of method signatures.
- **Decoupling**: Interfaces allow you to define dependencies on behavior rather than implementation.
- **Duck Typing**: "If it walks like a duck and quacks like a duck, it's a duck."
