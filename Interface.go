package main

import "fmt"

// I didn't add this in previous lesson "Composite Types"
// Struct can have methods, something like classes in other programming languages
// Basic a method is a function that belongs to a struct

// * Interfaces in Go provide a way to specify the behavior of an object.
// * An interface is a collection of method signatures, and any type that
// * implements all the methods of an interface is said to satisfy or implement that interface.

// Define an interface named Shape
type Shape interface {
	Area() float64
}

// Define a struct named Circle that implements the Shape interface
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define a struct named Rectangle that implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func usingInterface() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Use a function that accepts any Shape
	printArea(circle)
	printArea(rectangle)
}

// Function that accepts any type implementing the Shape interface
func printArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}

// * We define an interface Shape with a single method Area() that returns a float64.
// * We create two structs, Circle and Rectangle, and implement the Area() method for each. Since both structs have an Area() method, they implicitly satisfy the Shape interface.
// * The printArea function takes any type that implements the Shape interface and prints its area.
// * Interfaces allow you to write functions or methods that can work with different types as long as they adhere to a specific set of behaviors defined by the interface. This promotes flexibility and code reuse.
