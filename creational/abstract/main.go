package main

import "fmt"

// Define an interface with abstract methods
type Shape interface {
	Area() float64
}

// Create concrete types that implement the Shape interface
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5.0}
	rectangle := Rectangle{Width: 4.0, Height: 6.0}

	// Use the Shape interface to calculate areas
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
	}
}
