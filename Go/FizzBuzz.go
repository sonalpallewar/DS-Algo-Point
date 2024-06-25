package main

import (
    "fmt"
)

// Define a struct type for a basic shape
type Shape struct {
    Name string
}

// Method associated with the Shape struct
func (s *Shape) Describe() string {
    return fmt.Sprintf("This is a %s.", s.Name)
}

// Define a struct type for a Circle, embedding Shape
type Circle struct {
    Shape
    Radius float64
}

// Method associated with the Circle struct to calculate area
func (c *Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// Define a struct type for a Rectangle, embedding Shape
type Rectangle struct {
    Shape
    Width  float64
    Height float64
}

// Method associated with the Rectangle struct to calculate area
func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Function to demonstrate polymorphism using interfaces
type AreaCalculator interface {
    Area() float64
}

func PrintArea(a AreaCalculator) {
    fmt.Printf("The area is: %.2f\n", a.Area())
}

func main() {
    // Create an instance of Circle
    circle := Circle{
        Shape: Shape{
            Name: "Circle",
        },
        Radius: 5.0,
    }

    // Create an instance of Rectangle
    rectangle := Rectangle{
        Shape: Shape{
            Name: "Rectangle",
        },
        Width:  4.0,
        Height: 3.0,
    }

    // Print descriptions
    fmt.Println(circle.Describe())
    fmt.Println(rectangle.Describe())

    // Print areas using the PrintArea function
    PrintArea(&circle)
    PrintArea(&rectangle)
}
