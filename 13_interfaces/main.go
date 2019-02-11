package main

import (
	"fmt"
	"math"
)

// Shape definition
type Shape interface {
	area() float64
}

// Circle definition
type Circle struct {
	x, y, radius float64
}

// Rectangle definition
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{4, 5}
	fmt.Printf("%f, %f \n", getArea(circle), getArea(rectangle))
}
