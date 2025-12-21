# Go Data Types Program

This Go program demonstrates the **basic data types in Go**, including integers, floats, booleans, strings, runes (characters), and bytes.

---

## 📂 File

- `types.go` – Main program file demonstrating Go data types

---

## 🧾 Program Code

```go
package main

import "fmt"

func main() {
	// Integer types
	var a int = 10
	var b int8 = 100
	var c uint = 20

	// Floating point types
	var x float32 = 3.14
	var y float64 = 9.81

	// Boolean type
	var isGoFun bool = true

	// String type
	var name string = "Golang"

	// Character (rune)
	var ch rune = 'A' // rune is alias for int32

	// Byte
	var bt byte = 255 // byte is alias for uint8

	fmt.Println("Integer:", a, b, c)
	fmt.Println("Float:", x, y)
	fmt.Println("Boolean:", isGoFun)
	fmt.Println("String:", name)
	fmt.Println("Rune (char):", ch)
	fmt.Println("Byte:", bt)
}
```
