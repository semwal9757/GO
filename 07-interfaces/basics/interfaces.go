package main

import (
	"fmt"
	"math"
)

// Shape interface defines methods that all shapes must implement
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Describer interface for objects that can describe themselves
type Describer interface {
	Describe() string
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Describe returns a description of the rectangle
func (r Rectangle) Describe() string {
	return fmt.Sprintf("Rectangle with width %.2f and height %.2f", r.Width, r.Height)
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter (circumference) of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Describe returns a description of the circle
func (c Circle) Describe() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
	Side3  float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Side3
}

// printShapeInfo accepts any type that implements the Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// printDescription accepts any type that implements the Describer interface
func printDescription(d Describer) {
	fmt.Println(d.Describe())
}

// getTotalArea calculates total area of multiple shapes
func getTotalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	fmt.Println("--- Interface Basics ---")

	// Create different shapes
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{Base: 6, Height: 8, Side1: 6, Side2: 8, Side3: 10}

	// Using interface to work with different types
	fmt.Println("\n1. Rectangle:")
	printShapeInfo(rect)

	fmt.Println("\n2. Circle:")
	printShapeInfo(circle)

	fmt.Println("\n3. Triangle:")
	printShapeInfo(triangle)

	// Interface can hold any type that implements it
	var shape Shape
	shape = rect
	fmt.Printf("\nShape variable holding Rectangle - Area: %.2f\n", shape.Area())

	shape = circle
	fmt.Printf("Shape variable holding Circle - Area: %.2f\n", shape.Area())

	// Multiple interfaces
	fmt.Println("\n--- Multiple Interfaces ---")
	printDescription(rect)
	printDescription(circle)

	// Slice of interfaces
	fmt.Println("\n--- Slice of Interfaces ---")
	shapes := []Shape{rect, circle, triangle}
	for i, s := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n", i+1, s.Area(), s.Perimeter())
	}

	// Variadic function with interfaces
	fmt.Printf("\nTotal area of all shapes: %.2f\n", getTotalArea(rect, circle, triangle))

	// Empty interface can hold any value
	fmt.Println("\n--- Empty Interface ---")
	var anything interface{}
	anything = 42
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	anything = "Hello, Go!"
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)

	anything = rect
	fmt.Printf("Value: %v, Type: %T\n", anything, anything)
}
