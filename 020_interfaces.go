package main

import "fmt"
import "math"

type geometry interface { // basic interface for shapes
	area() float64
	perim() float64
}

// we'll implement this interface on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type
// we can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// the circle and the rect struct both implement
	// the geometry interface
	measure(r)
	measure(c)
}
