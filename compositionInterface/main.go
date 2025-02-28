// Go doesn't support inheritance but allows us to use composition for similar
// purposes
// Composition over inheritance in object-oriented programming (OOP) is the
// principle that classes should achieve polymorphic behavior and code
// reuse by their composition (by containing instances of other classes that
// implement the desired functionality) rather than inheritance
// from a base or parent class
package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Geometry interface {
	Area() float64
	Perimeter() float64
}

// The following methods implements the ones defining in the `Geometry` interface,
// this means `Circle` and `Rectangle` implements/satisfies the interface

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	figures := []Geometry{Circle{12}, Rectangle{2, 3}}
	for _, figure := range figures {
		fmt.Println(figure.Area())
	}
}
