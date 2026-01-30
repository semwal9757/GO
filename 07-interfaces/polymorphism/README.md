# Polymorphism

Polymorphism allows you to write code that can work with different types of objects, as long as they adhere to the same interface.

## 📌 How it Works

Because Go interfaces are implemented implicitly, you can create a function that accepts an interface type, and then pass any concrete type that implements that interface to the function.

```go
type Payer interface {
    Pay(amount float64)
}

func ProcessPayment(p Payer, amount float64) {
    p.Pay(amount)
}

// You can pass CreditCard, PayPal, or any other type that implements Pay()
```

## 🔑 Key Concepts

- **Flexibility**: Write generic code that handles multiple types.
- **Extensibility**: Add new types without changing existing code (Open/Closed Principle).
- **Collections**: Store different types in the same slice if they share an interface.
