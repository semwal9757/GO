# 07 - Interfaces and Polymorphism

Welcome to the **Interfaces** module! This section covers one of Go's most powerful features: interfaces. Interfaces enable polymorphism, code reusability, and clean abstractions in Go.

## � Module Structure

This module is organized into the following topics:

### 1. [Basics](./basics)
**`basics/interfaces.go`**
- Defining interfaces
- Implicit implementation
- Geometric shapes example

### 2. [Polymorphism](./polymorphism)
**`polymorphism/polymorphism.go`**
- Real-world polymorphism examples
- Payment & Notification systems
- Decoupling implementation from logic

### 3. [The Empty Interface](./empty-interface)
**`empty-interface/empty-interface.go`**
- Working with `interface{}` and `any`
- Generic containers
- Handling unknown types

### 4. [Type Assertions](./type-assertions)
**`type-assertions/type-assertions.go`**
- Extracting concrete types
- Safe vs. Unsafe assertions
- Type switches

## 🚀 Running the Examples

You can run each example by navigating to its directory:

```bash
# Basics
go run basics/interfaces.go

# Polymorphism
go run polymorphism/polymorphism.go

# Empty Interface
go run empty-interface/empty-interface.go

# Type Assertions
go run type-assertions/type-assertions.go
```

## 📖 Key Takeaways

1.  **Implicit Implementation**: Types implement interfaces automatically by implementing their methods.
2.  **Composition**: Interfaces can be composed of other interfaces.
3.  **Flexibility**: Accept interfaces in your functions to make them more testable and reusable.
4.  **Safety**: Always use the "comma ok" idiom (`val, ok := i.(Type)`) when performing type assertions.
